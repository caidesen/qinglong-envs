package router

import "net/http"

type Router interface {
	http.Handler

	// Group creates a new group with the same middleware stack as the original on top of the existing bundle.
	Group() Router
	// Route allows for configuring the Group inside the configureFn function.
	Route(configureFn func(Router))

	// Use adds middleware(s) to the Group.
	Use(middleware func(http.Handler) http.Handler, more ...func(http.Handler) http.Handler)
	// With adds new middleware(s) to the Group and returns a new Group with the updated middleware stack.
	With(middleware func(http.Handler) http.Handler, more ...func(http.Handler) http.Handler) Router

	// Mount creates a new group with a specified base path.
	Mount(basePath string) Router

	// Handle adds a new route to the Group's mux, applying all middlewares to the handler.
	Handle(basePath string, handler http.Handler)
	// HandleFunc registers the handler function for the given pattern to the Group's mux.
	// The handler is wrapped with the Group's middlewares.
	HandleFunc(pattern string, handler http.HandlerFunc)

	// Handler returns the handler and the pattern that matches the request.
	// It always returns a non-nil handler, see http.ServeMux.Handler documentation for details.
	Handler(r *http.Request) (h http.Handler, pattern string)

	// DisableNotFoundHandler disables the automatic registration of a not found handler for the root path.
	DisableNotFoundHandler()
	// NotFoundHandler sets a custom handler for the root path if no / route is registered.
	NotFoundHandler(handler http.HandlerFunc)
}
