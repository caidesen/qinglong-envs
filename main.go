package main

import (
	"fmt"
	"net/http"
	_ "qinglong-envs/pkg/env"
	"qinglong-envs/pkg/server"
)

func main() {
	conf := server.LoadHttpConfigFormEnv()
	httpServer := server.New(conf)
	httpServer.Register(func(router server.Router) {
		router.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("pong"))
		})
		router.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
			panic(fmt.Errorf("error"))
		})
	})
	httpServer.Start()
}
