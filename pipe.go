package pipe

import (
	"context"
	"fmt"
	"sync"

	"github.com/kubernetes-sigs/yaml"
	"github.com/wzshiming/pipe/components/common/load"
	"github.com/wzshiming/pipe/components/once"
	"github.com/wzshiming/pipe/components/stdio/input/inline"
	"golang.org/x/sync/errgroup"
)

type Pipe struct {
	config string
	group  *errgroup.Group
	ctx    context.Context
	cancel func()
	once   once.Once
	mut    sync.Mutex
}

type pipeCtxKeyType int

func GetPipeWithContext(ctx context.Context) (*Pipe, bool) {
	i := ctx.Value(pipeCtxKeyType(0))
	if i == nil {
		return nil, false
	}
	p, ok := i.(*Pipe)
	return p, ok
}

func NewPipeWithConfig(ctx context.Context, config []byte) (*Pipe, error) {
	config, err := yaml.YAMLToJSONStrict(config)
	if err != nil {
		return nil, err
	}
	conf := string(config)
	c := &Pipe{}
	c.group, c.ctx = errgroup.WithContext(ctx)
	c.ctx = context.WithValue(c.ctx, pipeCtxKeyType(0), c)
	var o once.Once
	err = load.Load(c.ctx, inline.NewInlineWithConfig(&inline.Config{Data: conf}), &o)
	if err != nil {
		return nil, err
	}
	if o == nil {
		return nil, fmt.Errorf("no entry")
	}
	c.once = o
	c.config = conf
	return c, nil
}

func (c *Pipe) Start() error {
	c.mut.Lock()
	defer c.mut.Unlock()
	err := c.start(c.once)
	if err != nil {
		return err
	}
	return nil
}

func (c *Pipe) Run() error {
	err := c.Start()
	if err != nil {
		return err
	}
	return c.Wait()
}

func (c *Pipe) Wait() error {
	return c.group.Wait()
}

func (c *Pipe) start(o once.Once) error {
	ctx, cancel := context.WithCancel(c.ctx)
	run := func() error {
		return o.Do(ctx)
	}
	c.group.Go(run)
	if c.cancel != nil {
		c.cancel()
	}
	c.once = o
	c.cancel = cancel
	return nil
}

func (c *Pipe) Reload(config []byte) error {
	config, err := yaml.YAMLToJSONStrict(config)
	if err != nil {
		return err
	}
	conf := string(config)
	var o once.Once
	err = load.Load(c.ctx, inline.NewInlineWithConfig(&inline.Config{Data: conf}), &o)
	if err != nil {
		return err
	}
	if o == nil {
		return fmt.Errorf("no entry")
	}

	c.mut.Lock()
	defer c.mut.Unlock()
	err = c.start(o)
	if err != nil {
		return err
	}
	c.config = conf
	return nil
}

func (c *Pipe) Close() error {
	c.mut.Lock()
	defer c.mut.Unlock()
	if c.cancel != nil {
		c.cancel()
	}
	return nil
}

func (c *Pipe) Config() []byte {
	c.mut.Lock()
	defer c.mut.Unlock()
	return []byte(c.config)
}
