package ref

import (
	"context"
	"fmt"

	"github.com/wzshiming/pipe/components"
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/stream"
)

var (
	ErrNotHandler = fmt.Errorf("not handler")
)

const name = "ref"

func init() {
	configure.Register(name, NewRefWithConfig)
}

type Config struct {
	Ref string
}

func NewRefWithConfig(ctx context.Context, conf *Config) (stream.Handler, error) {
	components, ok := components.GetCtxComponents(ctx)
	if !ok || components == nil || components.StreamHandlers == nil {
		return nil, ErrNotHandler
	}
	handler, ok := components.StreamHandlers[conf.Ref]
	if !ok {
		return nil, fmt.Errorf("%s: %w", conf.Ref, ErrNotHandler)
	}
	return handler, nil
}