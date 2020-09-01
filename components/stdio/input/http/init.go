package http

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/wzshiming/pipe/components/common/register"
	"github.com/wzshiming/pipe/components/stdio/input"
	"github.com/wzshiming/pipe/components/stream"
	"github.com/wzshiming/pipe/internal/round_tripper"
)

const (
	name = "http"
)

func init() {
	register.Register(name, NewHTTPWithConfig)
}

type Config struct {
	Dialer stream.Dialer `json:",omitempty"`
	URL    string
}

func NewHTTPWithConfig(conf *Config) (input.Input, error) {
	cli := http.Client{
		Transport: round_tripper.RoundTripper(conf.Dialer),
	}
	resp, err := cli.Get(conf.URL)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return bytes.NewBuffer(body), nil
}
