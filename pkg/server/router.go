package server

import (
	"net/http"
	"slices"
)

type Router interface {
	http.Handler
	Use(mws ...Middleware)
	Group(func(r *Mux))

	Get(path string, fn http.HandlerFunc, mx ...Middleware)
	Post(path string, fn http.HandlerFunc, mx ...Middleware)
	Put(path string, fn http.HandlerFunc, mx ...Middleware)
	Patch(path string, fn http.HandlerFunc, mx ...Middleware)
	Delete(path string, fn http.HandlerFunc, mx ...Middleware)
	Options(path string, fn http.HandlerFunc, mx ...Middleware)
	Head(path string, fn http.HandlerFunc, mx ...Middleware)

	Mount(path string, sub *Mux)
}

type (
	Middleware func(http.Handler) http.Handler
	Mux        struct {
		*http.ServeMux
		chain []Middleware
	}
)

func NewRouter() Router {
	return &Mux{ServeMux: &http.ServeMux{}}
}

func (r *Mux) Use(mws ...Middleware) {
	r.chain = append(r.chain, mws...)
}

func (r *Mux) Group(fn func(r *Mux)) {
	fn(&Mux{ServeMux: r.ServeMux, chain: slices.Clone(r.chain)})
}

func (r *Mux) Get(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodGet, path, fn, mx)
}

func (r *Mux) Post(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodPost, path, fn, mx)
}

func (r *Mux) Put(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodPut, path, fn, mx)
}

func (r *Mux) Delete(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodDelete, path, fn, mx)
}

func (r *Mux) Head(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodHead, path, fn, mx)
}

func (r *Mux) Options(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodOptions, path, fn, mx)
}

func (r *Mux) Patch(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodPatch, path, fn, mx)
}

func (r *Mux) Mount(path string, sub *Mux) {
	r.Handle(path, sub)
}

func (r *Mux) handle(method, path string, handler http.HandlerFunc, mx []Middleware) {
	r.Handle(method+" "+path, r.chainWrap(handler, mx))
}

func (r *Mux) chainWrap(fn http.Handler, mx []Middleware) http.Handler {
	if len(mx) == 0 {
		return fn
	}
	out := fn
	mx = append(slices.Clone(r.chain), mx...)

	slices.Reverse(mx)

	for _, m := range mx {
		out = m(out)
	}

	return out
}
