package router

import "net/http"

type Middleware func(http.Handler) http.Handler
type Router interface {
	http.Handler
	Use(mws ...Middleware)
	Group(func(group Router))
	GroupWithPrefix(prefix string, fn func(Router))

	Get(path string, fn http.Handler, mx ...Middleware)
	Post(path string, fn http.Handler, mx ...Middleware)
	Put(path string, fn http.Handler, mx ...Middleware)
	Patch(path string, fn http.Handler, mx ...Middleware)
	Delete(path string, fn http.Handler, mx ...Middleware)
	Options(path string, fn http.Handler, mx ...Middleware)
	Head(path string, fn http.Handler, mx ...Middleware)

	Mount(path string, sub Router)
}
