package forward

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/pipeproxy/pipe/internal/http/template"
	"github.com/pipeproxy/pipe/internal/pool"
)

type Forward struct {
	url       template.Format
	transport http.RoundTripper
}

func NewForward(url string, transport http.RoundTripper) (*Forward, error) {
	f := &Forward{
		transport: transport,
	}
	if url == "" {
		url = "{{.Scheme}}://{{.Host}}"
	}

	u, err := template.NewFormat(url)
	if err != nil {
		return nil, err
	}
	f.url = u
	return f, nil
}

func (h *Forward) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	proxy := httputil.ReverseProxy{
		BufferPool:   pool.Bytes,
		Transport:    h.transport,
		ErrorHandler: errorHandler,
	}
	u := h.url.FormatString(r)
	target, err := url.Parse(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	proxy.Director = directorFunc(target)
	proxy.ServeHTTP(rw, r)
}

func directorFunc(target *url.URL) func(req *http.Request) {
	return func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if target.RawQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = target.RawQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = target.RawQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
}

func errorHandler(rw http.ResponseWriter, r *http.Request, err error) {
	http.Error(rw, err.Error(), http.StatusInternalServerError)
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
