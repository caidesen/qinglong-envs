package server

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
	"qinglong-envs/internal/httperr"
	"strconv"
)

type Handler struct {
	validate *validator.Validate
}

func NewHandler(validate *validator.Validate) *Handler {
	return &Handler{validate: validate}
}

func (h *Handler) WriteJSON(w http.ResponseWriter, resp any) {
	if resp == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) WriteErr(w http.ResponseWriter, err error) {
	var httpErr *httperr.HTTPError
	if !errors.As(err, &httpErr) {
		httpErr = httperr.InternalServerError(err.Error())
	}
	httpErr.Out(w)
}

func (h *Handler) GetIntInPath(r *http.Request, name string) (int, error) {
	pathId := r.PathValue(name)
	id, err := strconv.Atoi(pathId)
	if err != nil {
		return 0, httperr.InputError("invalid id")
	}
	return id, nil
}

func (h *Handler) BindJSON(r *http.Request, i any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return httperr.InputError("invalid content type")
	}
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		return httperr.InputError("invalid json body")
	}
	return nil
}

func (h *Handler) Validate(s any) error {
	err := h.validate.Struct(s)
	if err != nil {
		return httperr.InputError("input validation error").WithDetail(err.Error())
	}
	return err
}

type PaginationParams struct {
	PageSize int `query:""`
	Current  int `query:"current"`
}

func (p *PaginationParams) Offset() int {
	return (p.Current - 1) * p.PageSize
}

// PaginationResult PaginationResult[T]
//
//	@Description	分页查询结果
type PaginationResult[T any] struct {
	Current int  `json:"current"`
	Total   int  `json:"total"`
	List    []*T `json:"list"`
}

func NewPaginationResult[T any](current, total int, list []*T) *PaginationResult[T] {
	return &PaginationResult[T]{Current: current, Total: total, List: list}
}
