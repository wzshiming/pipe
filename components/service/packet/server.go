package packet

import (
	"context"
	"sync"

	"github.com/pipeproxy/pipe/components/packet"
	"github.com/pipeproxy/pipe/internal/listener"
	"github.com/wzshiming/logger"
)

type Server struct {
	listenConfig packet.ListenConfig
	packet       packet.Packet
	handler      packet.Handler
	pool         sync.Map
}

func NewServer(listenConfig packet.ListenConfig, handler packet.Handler) (*Server, error) {
	s := &Server{
		listenConfig: listenConfig,
		handler:      handler,
	}

	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	pkt, err := s.listenConfig.ListenPacket(ctx)
	if err != nil {
		return err
	}
	s.packet = pkt
	s.ServePacket(ctx, pkt)
	return nil
}

func (s *Server) Close() error {
	if s.packet == nil {
		return nil
	}
	err := s.packet.Close()
	if listener.IsClosedConnError(err) {
		err = nil
	}
	return err
}

func (s *Server) ServePacket(ctx context.Context, pkt packet.Packet) {
	log := logger.FromContext(ctx)
	log = log.WithName("packet")
	ctx = logger.WithContext(ctx, log)
	s.handler.ServePacket(ctx, nopCloser{pkt})
	err := pkt.Close()
	if listener.IsClosedConnError(err) {
		err = nil
	}
	if err != nil {
		log.Error(err, "close listen")
		return
	}
}

type nopCloser struct {
	packet.Packet
}

func (nopCloser) Close() error {
	return nil
}
