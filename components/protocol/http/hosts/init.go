package hosts

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
)

const (
	name = "hosts"
)

func init() {
	register.Register(name, NewHostsWithConfig)
}

var (
	ErrNotHandler = fmt.Errorf("error not handler")
	ErrNotDomain  = fmt.Errorf("error not router")
)

type Route struct {
	Domains []string
	Handler http.Handler
}

type Config struct {
	Hosts    []*Route
	NotFound http.Handler `json:",omitempty"`
}

func NewHostsWithConfig(conf *Config) (http.Handler, error) {
	mux := NewHosts()
	mux.NotFound(conf.NotFound)
	for _, route := range conf.Hosts {
		if route.Handler == nil {
			return nil, ErrNotHandler
		}
		if len(route.Domains) == 0 {
			return nil, ErrNotDomain
		}
		for _, domain := range route.Domains {
			if domain == "" {
				return nil, ErrNotDomain
			}
			err := mux.Handle(domain, route.Handler)
			if err != nil {
				return nil, err
			}
		}
	}
	return mux, nil
}
