package tls

import (
	"context"

	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/components/tls"
	"github.com/wzshiming/logger"
)

type Tls struct {
	listenConfig stream.ListenConfig
	tlsConfig    tls.TLS
}

func NewTls(listenConfig stream.ListenConfig, tlsConfig tls.TLS) *Tls {
	return &Tls{
		listenConfig: listenConfig,
		tlsConfig:    tlsConfig,
	}
}

func (d *Tls) ListenStream(ctx context.Context) (stream.StreamListener, error) {
	log := logger.FromContext(ctx)
	log = log.WithName("tls")
	ctx = logger.WithContext(ctx, log)
	listener, err := d.listenConfig.ListenStream(ctx)
	if err != nil {
		return nil, err
	}
	return tls.NewListener(listener, d.tlsConfig.TLS()), nil
}
