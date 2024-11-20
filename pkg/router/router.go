package router

import "net/http"

type Middleware func(http.Handler) http.Handler
type Router interface {
	http.Handler
	Use(mws ...Middleware)
	Group(func(group Router))

	Get(path string, fn http.HandlerFunc, mx ...Middleware)
	Post(path string, fn http.HandlerFunc, mx ...Middleware)
	Put(path string, fn http.HandlerFunc, mx ...Middleware)
	Patch(path string, fn http.HandlerFunc, mx ...Middleware)
	Delete(path string, fn http.HandlerFunc, mx ...Middleware)
	Options(path string, fn http.HandlerFunc, mx ...Middleware)
	Head(path string, fn http.HandlerFunc, mx ...Middleware)

	Mount(path string, sub Router)
}
