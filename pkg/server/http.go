package server

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

func LoadListeningFormEnv() string {
	port := 3000
	host := "127.0.0.1"
	portEnv := os.Getenv("QINGLONG_ENVS_SERVER_PORT")
	hostEnv := os.Getenv("QINGLONG_ENVS_SERVER_HOST")
	// to int
	portEnvInt, _ := strconv.Atoi(portEnv)
	if portEnvInt != 0 {
		port = portEnvInt
	}
	if hostEnv != "" {
		host = hostEnv
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func StartHttpServer(h http.Handler) {
	addr := LoadListeningFormEnv()
	server := &http.Server{
		Addr:    addr,
		Handler: h,
	}
	slog.Info(fmt.Sprintf("listening on %s", addr))
	log.Fatal(server.ListenAndServe())
}
