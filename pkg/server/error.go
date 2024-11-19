package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPError struct {
	Code     int
	Message  interface{}
	Internal error `json:"-"` // Stores the error returned by an external dependency
}

// Error makes it compatible with `error` interface.
func (he *HTTPError) Error() string {
	if he.Internal == nil {
		return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
	}
	return fmt.Sprintf("code=%d, message=%v, internal=%v", he.Code, he.Message, he.Internal)
}

func NewHttpError(code int, message string) error {
	return &HTTPError{
		Code:    code,
		Message: message,
	}
}

type ErrorResponse struct {
	Msg     string `json:"msg"`
	ErrCode int    `json:"errCode"`
}

func ErrorHandler(w http.ResponseWriter, err error) {
	httpStatus := http.StatusInternalServerError
	resp := ErrorResponse{
		Msg:     err.Error(),
		ErrCode: -1,
	}
	httpErr, ok := err.(*HTTPError)
	if ok {
		httpStatus = httpErr.Code
	}
	w.WriteHeader(httpStatus)
	marshal, _ := json.Marshal(resp)
	w.Write(marshal)
}
