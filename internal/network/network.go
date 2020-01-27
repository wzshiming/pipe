package network

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

var (
	mut   = sync.Mutex{}
	cache = map[string]*Hub{}
)

func CloseExcess() {
	mut.Lock()
	defer mut.Unlock()
	dk := []string{}
	for k, v := range cache {
		if v.size == 0 {
			v.Close()
			dk = append(dk, k)
		}
	}

	for _, k := range dk {
		log.Printf("[INFO] Close listen to %s", k)
		delete(cache, k)
	}
}

func Listen(network, address string) (net.Listener, error) {
	mut.Lock()
	defer mut.Unlock()
	key := fmt.Sprintf("%s://%s", network, address)
	n, ok := cache[key]
	if ok {
		log.Printf("[INFO] Relisten to %s", key)
		return n.Listener(), nil
	}

	log.Printf("[INFO] Listen to %s", key)
	l, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	n = newListener(l)
	cache[key] = n
	return n.Listener(), nil
}

type Hub struct {
	listener net.Listener
	ch       chan net.Conn
	size     int32
}

func newListener(listener net.Listener) *Hub {
	m := &Hub{
		listener: listener,
		ch:       make(chan net.Conn),
	}
	go m.run()
	return m
}

func (h *Hub) run() {
	for {
		conn, err := h.listener.Accept()
		if err != nil {
			return
		}
		h.ch <- conn
	}
}

func (h *Hub) Close() error {
	h.listener.Close()
	close(h.ch)

	return nil
}

func (h *Hub) Listener() *Listener {
	l := &Listener{
		hub:  h,
		ch:   h.ch,
		exit: make(chan struct{}),
	}
	h.size++
	return l
}

type Listener struct {
	closeOnce sync.Once
	ch        chan net.Conn
	exit      chan struct{}
	hub       *Hub
}

func (l *Listener) Accept() (net.Conn, error) {
	select {
	case conn := <-l.ch:
		return conn, nil
	case <-l.exit:
		return nil, io.ErrClosedPipe
	}
}

// Close closes the listener.
func (l *Listener) Close() error {
	l.closeOnce.Do(func() {
		l.hub.size--
		l.exit <- struct{}{}
	})
	return nil
}

// Addr returns the listener's network address.
func (l *Listener) Addr() net.Addr {
	return l.hub.listener.Addr()
}