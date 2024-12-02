package httpapi

import "net/http"

type Middleware func(http.Handler) http.Handler
type Router interface {
	http.Handler
	Use(mws ...Middleware)
	Group(func(group Router))

	Handle(path string, handler http.Handler, mx ...Middleware)
	API(path string, handler HandlerFunc, mx ...Middleware)
}
