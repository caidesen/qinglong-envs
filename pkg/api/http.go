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

func GetIntInPath(r *http.Request, name string) (int, error) {
	pathId := r.PathValue(name)
	id, err := strconv.Atoi(pathId)
	if err != nil {
		return 0, ValidateError("invalid id")
	}
	return id, nil
}

func BindJSONBody[T any](r *http.Request) (T, error) {
	var i T
	if r.Header.Get("Content-Type") != "application/json" {
		return i, InvalidContentTypeError
	}
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return i, InvalidJSONError
	}
	return i, nil
}

func Response(w http.ResponseWriter, resp any) error {
	if resp == nil {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func ErrorHandler(w http.ResponseWriter, err error) {
	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		httpErr = InternalError(err.Error())
	}
	httpErr.Out(w)
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		ErrorHandler(w, err)
	} else {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusOK)
	}
}

type PaginationParams struct {
	PageSize int `json:"limit"`
	Current  int `json:"current"`
}

func (p *PaginationParams) Offset() int {
	return (p.Current - 1) * p.PageSize
}

type Pagination struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}
