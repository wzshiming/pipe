package decode

import (
	"context"
	"reflect"
	"testing"
)

type Config struct {
	Name string
}

func (c Config) M() {}

type Adapter interface {
	M()
}

func TestDecodeStruct(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx    context.Context
		config []byte
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			args: args{ctx, []byte(`{"@name":"hello"}`)},
			want: &Config{"hello"},
		},
		{
			args: args{ctx, []byte(`[{"@name":"hello"},{"@name":"hello2"}]`)},
			want: []*Config{{"hello"}, {"hello2"}},
		},
		{
			args: args{ctx, []byte(`{"A":{"@name":"hello"}}`)},
			want: &struct{ A *Config }{&Config{"hello"}},
		},
		{
			args: args{ctx, []byte(`{"A":{"@name":"hello"},"B":[{"@name":"hello2"},{"@name":"hello3"}]}`)},
			want: &struct {
				A *Config
				B []*Config
			}{&Config{"hello"}, []*Config{{"hello2"}, {"hello3"}}},
		},

		{
			args: args{ctx, []byte(`{"@name":"hello"}`)},
			want: Config{"hello"},
		},
		{
			args: args{ctx, []byte(`[{"@name":"hello"},{"@name":"hello2"}]`)},
			want: []Config{{"hello"}, {"hello2"}},
		},
		{
			args: args{ctx, []byte(`{"A":{"@name":"hello"}}`)},
			want: &struct{ A Config }{Config{"hello"}},
		},
		{
			args: args{ctx, []byte(`{"A":{"@name":"hello"},"B":[{"@name":"hello2"},{"@name":"hello3"}]}`)},
			want: &struct {
				A Config
				B []Config
			}{Config{"hello"}, []Config{{"hello2"}, {"hello3"}}},
		},
	}

	fun := []interface{}{
		func(name string, config []byte) (*Config, error) {
			return &Config{name}, nil
		},
		func(name string, config []byte) (Adapter, error) {
			return &Config{Name: name}, nil
		},
	}

	for _, f := range fun {
		stdManager := newDecoderManager()
		stdDecoder := &decoder{
			decoderManager: stdManager,
		}

		err := stdManager.Register(f)
		if err != nil {
			t.Fatal(err)
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotValue := reflect.New(reflect.TypeOf(tt.want))
				if err := stdDecoder.Decode(tt.args.ctx, tt.args.config, gotValue.Interface()); (err != nil) != tt.wantErr {
					t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				}

				got := gotValue.Elem().Interface()
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Decode() got = %#v, want %#v", got, tt.want)
				}
			})
		}
	}
}
