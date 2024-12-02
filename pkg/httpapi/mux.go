package httpapi

import (
	"net/http"
	"slices"
)

type Mux struct {
	*http.ServeMux
	chain []Middleware
}

func NewRouter() Router {
	m := Mux{ServeMux: http.NewServeMux()}
	return &m
}

func (r *Mux) Use(mws ...Middleware) {
	r.chain = append(r.chain, mws...)
}

func (r *Mux) Group(fn func(Router)) {
	fn(&Mux{ServeMux: r.ServeMux, chain: slices.Clone(r.chain)})
}

func (r *Mux) Handle(path string, handler http.Handler, mx ...Middleware) {
	r.ServeMux.Handle(path, r.chainWrap(handler, mx))
}

func (r *Mux) API(path string, handler HandlerFunc, mx ...Middleware) {
	r.Handle(path, handler, mx...)
}

func (r *Mux) chainWrap(fn http.Handler, mx []Middleware) http.Handler {
	out := fn
	mx = append(slices.Clone(r.chain), mx...)
	if len(mx) == 0 {
		return fn
	}
	slices.Reverse(mx)
	for _, m := range mx {
		out = m(out)
	}

	return out
}
