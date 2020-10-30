package forward

import (
	"net/http"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stream"
	"github.com/pipeproxy/pipe/internal/round_tripper"
)

const (
	name = "forward"
)

func init() {
	register.Register(name, NewForwardWithConfig)
}

type Config struct {
	Dialer stream.Dialer `json:",omitempty"`
	URL    string
}

// NewForwardWithConfig create a new forward with config.
func NewForwardWithConfig(conf *Config) (http.Handler, error) {
	return NewForward(conf.URL, round_tripper.RoundTripper(conf.Dialer))
}
