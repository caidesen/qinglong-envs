package server

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
)

type HttpConfig struct {
	Port int
	Host string
}

func DefaultHttpConfig() HttpConfig {
	return HttpConfig{
		Port: 3000,
		Host: "0.0.0.0",
	}
}

func LoadHttpConfigFormEnv() HttpConfig {
	c := DefaultHttpConfig()
	portEnv := os.Getenv("QINGLONG_ENVS_SERVER_PORT")
	hostEnv := os.Getenv("QINGLONG_ENVS_SERVER_HOST")
	// to int
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		slog.Warn("load env QINGLONG_ENVS_SERVER_PORT: %s", err)
	} else {
		c.Port = port
	}
	ipOK := net.ParseIP(hostEnv)
	if ipOK != nil {
		c.Host = hostEnv
	}
	return c
}

type HttpServer struct {
	conf   HttpConfig
	router Router
}

func New(conf HttpConfig) HttpServer {
	r := NewRouter()
	r.Use(RecoverMiddleware)
	//r.Use(Middleware.Recoverer)
	return HttpServer{
		conf:   conf,
		router: r,
	}
}

func (s *HttpServer) Start() {
	addr := fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}
	slog.Info(fmt.Sprintf("listening on %s", addr))
	log.Fatal(server.ListenAndServe())
}

func (s *HttpServer) Register(fn func(Router)) {
	fn(s.router)
}
