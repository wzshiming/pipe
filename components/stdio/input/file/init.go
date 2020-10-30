package file

import (
	"bytes"
	"io/ioutil"

	"github.com/pipeproxy/pipe/components/common/register"
	"github.com/pipeproxy/pipe/components/stdio/input"
)

const (
	name = "file"
)

func init() {
	register.Register(name, NewFileWithConfig)
}

type Config struct {
	Path string
}

func NewFileWithConfig(conf *Config) input.Input {
	return input.NewLazyReader(func() (input.Input, error) {
		data, err := ioutil.ReadFile(conf.Path)
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(data), nil
	})
}
