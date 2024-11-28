package router

import (
	"net/http"
	"slices"
)

type Mux struct {
	*http.ServeMux
	chain  []Middleware
	prefix string
}

func NewRouter(prefix ...string) Router {
	m := Mux{ServeMux: &http.ServeMux{}}
	if len(prefix) > 0 {
		m.prefix = prefix[0]
	}
	return &m
}

func (r *Mux) Use(mws ...Middleware) {
	r.chain = append(r.chain, mws...)
}

func (r *Mux) Group(fn func(Router)) {
	fn(&Mux{ServeMux: r.ServeMux, chain: slices.Clone(r.chain), prefix: r.prefix})
}

func (r *Mux) GroupWithPrefix(prefix string, fn func(Router)) {
	fn(&Mux{ServeMux: r.ServeMux, chain: slices.Clone(r.chain), prefix: r.prefix + prefix})
}

func (r *Mux) Get(path string, fn http.Handler, mx ...Middleware) {
	r.handle(http.MethodGet, path, fn, mx)
}

func (r *Mux) Post(path string, fn http.Handler, mx ...Middleware) {
	r.handle(http.MethodPost, path, fn, mx)
}

func (r *Mux) Put(path string, fn http.Handler, mx ...Middleware) {
	r.handle(http.MethodPut, path, fn, mx)
}

func (r *Mux) Delete(path string, fn http.Handler, mx ...Middleware) {
	r.handle(http.MethodDelete, path, fn, mx)
}

func (r *Mux) Head(path string, fn http.Handler, mx ...Middleware) {
	r.handle(http.MethodHead, path, fn, mx)
}

func (r *Mux) Options(path string, fn http.Handler, mx ...Middleware) {
	r.handle(http.MethodOptions, path, fn, mx)
}

func (r *Mux) Patch(path string, fn http.Handler, mx ...Middleware) {
	r.handle(http.MethodPatch, path, fn, mx)
}

func (r *Mux) Mount(path string, sub Router) {
	r.Handle(path, http.StripPrefix(path, sub))
}

func (r *Mux) handle(method, path string, handler http.Handler, mx []Middleware) {
	r.Handle(method+" "+r.prefix+path, r.chainWrap(handler, mx))
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

type X struct {
	*http.ServeMux
	chain []Middleware
}

func NewX() *X {
	return &X{ServeMux: http.NewServeMux()}
}

func (r *X) Use(mws ...Middleware) {
	r.chain = append(r.chain, mws...)
}

func (r *X) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.ServeMux.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		r.chainWrap(http.HandlerFunc(handler)).ServeHTTP(writer, request)
	})
}

func (r *X) chainWrap(fn http.Handler) http.Handler {
	out := fn
	mx := slices.Clone(r.chain)
	if len(mx) == 0 {
		return fn
	}
	slices.Reverse(mx)
	for _, m := range mx {
		out = m(out)
	}

	return out
}
