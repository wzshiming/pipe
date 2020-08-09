package quic

import (
	"context"

	quic "github.com/lucas-clemente/quic-go"
	"github.com/wzshiming/pipe/components/packet"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/components/tls"
)

type server struct {
	pkt       packet.Packet
	tlsConfig *tls.Config
}

func NewServer(pkt packet.Packet, tlsConfig *tls.Config) stream.ListenConfig {
	s := &server{
		pkt:       pkt,
		tlsConfig: tlsConfig,
	}
	return s
}

func (s *server) ListenStream(ctx context.Context) (stream.StreamListener, error) {
	listen, err := quic.Listen(s.pkt, s.tlsConfig, &quic.Config{})
	if err != nil {
		return nil, err
	}
	return NewListener(ctx, listen), nil
}