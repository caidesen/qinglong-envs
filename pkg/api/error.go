package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPError struct {
	HttpStatus int    `json:"-"`
	Message    string `json:"message,omitempty"`
	Detail     any    `json:"detail,omitempty"`
}

// Error makes it compatible with `error` interface.
func (he *HTTPError) Error() string {
	return fmt.Sprintf("msg:%s", he.Message)
}

func (he *HTTPError) WithDetail(detail any) *HTTPError {
	he.Detail = detail
	return he
}

func (he *HTTPError) Out(w http.ResponseWriter) {
	type errResp struct {
		Error *HTTPError `json:"error"`
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(he.HttpStatus)
	err := json.NewEncoder(w).Encode(&errResp{Error: he})
	if err != nil {
		return
	}
}

func New(status int, message string) *HTTPError {
	e := &HTTPError{
		HttpStatus: status,
		Message:    message,
	}
	return e
}

func ValidateError(msg string) *HTTPError {
	return New(http.StatusBadRequest, msg)
}

func NotFoundError(msg string) *HTTPError {
	return New(http.StatusNotFound, msg)
}

func InternalError(msg string) *HTTPError {
	return New(http.StatusInternalServerError, msg)
}

func NotPermittedError(msg string) *HTTPError {
	return New(http.StatusForbidden, msg)
}

func NotAuthorizedError(msg string) *HTTPError {
	return New(http.StatusUnauthorized, msg)
}
