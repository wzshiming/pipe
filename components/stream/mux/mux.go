package mux

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"strings"
	"sync/atomic"

	"github.com/wzshiming/crun"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/trie"
)

var (
	ErrNotFound = fmt.Errorf("error not found")
)

// Mux is an Applicative protocol multiplexer
// It matches the prefix of each incoming reader against a list of registered patterns
// and calls the handler for the pattern that most closely matches the Handler.
type Mux struct {
	trie         *trie.Trie
	prefixLength int
	size         uint32
	handlers     map[uint32]stream.Handler
	notFound     stream.Handler
}

// NewProtoMux create a new Mux.
func NewMux() *Mux {
	p := &Mux{
		trie:     trie.NewTrie(),
		handlers: map[uint32]stream.Handler{},
	}

	return p
}

// NotFound replies to the handler with an Handler not found error.
func (m *Mux) NotFound(handler stream.Handler) error {
	m.notFound = handler
	return nil
}

func (m *Mux) HandleRegexp(pattern string, handler stream.Handler) error {
	if !strings.HasPrefix(pattern, "^") {
		return fmt.Errorf("only prefix matching is supported, change to %q", "^"+pattern)
	}
	r, err := crun.Compile(pattern)
	if err != nil {
		return err
	}

	if size := r.Size(); size > 1000 {
		return fmt.Errorf("regular is too large: %d", size)
	}

	buf := m.setHandler(handler)
	r.Range(func(prefix string) bool {
		m.handle(prefix, buf)
		return true
	})
	return nil
}

func (m *Mux) HandlePrefix(prefix string, handler stream.Handler) error {
	buf := m.setHandler(handler)
	m.handle(prefix, buf)
	return nil
}

// Handler returns most matching handler and prefix bytes data to use for the given reader.
func (m *Mux) Handler(r io.Reader) (handler stream.Handler, prefix []byte, err error) {
	if m.prefixLength == 0 {
		return nil, nil, ErrNotFound
	}
	parent := m.trie.Mapping()
	off := 0
	prefix = make([]byte, m.prefixLength)
	for {
		i, err := r.Read(prefix[off:])
		if err != nil {
			return nil, nil, err
		}
		if i == 0 {
			break
		}

		data, next, _ := parent.Get(prefix[off : off+i])
		if len(data) != 0 {
			conn, ok := m.getHandler(data)
			if ok {
				handler = conn
			}
		}

		off += i
		if next == nil {
			break
		}
		parent = next
	}

	if handler == nil {
		if m.notFound == nil {
			return nil, prefix[:off], ErrNotFound
		}
		handler = m.notFound
	}
	return handler, prefix[:off], nil
}

func (m *Mux) handle(prefix string, buf []byte) {
	m.trie.Put([]byte(prefix), buf)
	if m.prefixLength < len(prefix) {
		m.prefixLength = len(prefix)
	}
}

func (m *Mux) setHandler(hand stream.Handler) []byte {
	k := atomic.AddUint32(&m.size, 1)
	m.handlers[k] = hand
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, k)
	return buf
}

func (m *Mux) getHandler(index []byte) (stream.Handler, bool) {
	c, ok := m.handlers[binary.BigEndian.Uint32(index)]
	return c, ok
}

// ServeStream dispatches the reader to the handler whose pattern most closely matches the reader.
func (m *Mux) ServeStream(ctx context.Context, stm stream.Stream) {
	connector, buf, err := m.Handler(stm)
	if err != nil {
		log.Printf("[ERROR] prefix %q: %s", buf, err.Error())
		stm.Close()
		return
	}
	stm = UnreadStream(stm, buf)
	connector.ServeStream(ctx, stm)
}