package main

import (
	"net/http"
	_ "qinglong-envs/pkg/env"
	"qinglong-envs/pkg/middleware"
	"qinglong-envs/pkg/router"
	"qinglong-envs/pkg/server"
)

func main() {
	r := router.NewRouter()
	r.Use(middleware.RecoverMiddleware)
	r.Group(func(g router.Router) {
		g.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("pong"))
		})
	})
	server.StartHttpServer(r)
}
