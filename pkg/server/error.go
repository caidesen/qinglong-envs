package server

import (
	"fmt"
	"net/http"
)

// RFC9457 https://www.rfc-editor.org/rfc/rfc9457.html
type HttpErrorDetail struct {
	Detail  string `json:"detail,omitempty"`
	Pointer string `json:"pointer,omitempty"`
}

type HTTPError struct {
	Type     string            `json:"type"`
	Title    string            `json:"title"`
	Status   int               `json:"status,omitempty"`
	Detail   string            `json:"detail,omitempty"`
	Instance string            `json:"instance,omitempty"`
	Errors   []HttpErrorDetail `json:"errors,omitempty"`
}

// Error makes it compatible with `error` interface.
func (he *HTTPError) Error() string {
	if he.Title == "" && he.Detail == "" {
		return fmt.Sprintf("Status %d", he.Status)
	}

	if he.Detail == "" {
		return fmt.Sprintf("%s", he.Title)
	}
	return fmt.Sprintf("%s: %s", he.Title, he.Detail)
}

func New(statusCode int, problemType, title, detail, instance string, errors []HttpErrorDetail) *HTTPError {
	if problemType == "" {
		problemType = "about:blank"
	}
	if problemType == "about:blank" {
		title = http.StatusText(statusCode)
	}
	return &HTTPError{
		Type:     problemType,
		Title:    title,
		Status:   statusCode,
		Detail:   detail,
		Instance: instance,
		Errors:   errors,
	}
}
