package api

import (
	"encoding/json"
	"errors"
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

var InvalidJSONError = ValidateError("invalid json body")
var InvalidContentTypeError = ValidateError("invalid content type")

func BindJSONBody(i any, r *http.Request) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return InvalidContentTypeError
	}
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		return InvalidJSONError
	}
	return nil
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		var httpErr *HTTPError
		if errors.As(err, &httpErr) {
			err.(*HTTPError).Out(w)
		} else {
			InternalError(err.Error()).Out(w)
		}
	}
}

func PostApi() {

}
