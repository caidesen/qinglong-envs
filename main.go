package main

import (
	_ "qinglong-envs/pkg/env"
	"qinglong-envs/pkg/middleware"
	"qinglong-envs/pkg/router"
	"qinglong-envs/pkg/server"
)

func main() {
	r := router.NewRouter()
	r.Use(middleware.Recovery)
	server.StartHttpServer(r)
}
