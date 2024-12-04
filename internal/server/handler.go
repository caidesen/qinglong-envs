package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"qinglong-envs/internal/httperr"
	"strconv"
)

func WriteJSON(w http.ResponseWriter, resp any) {
	if resp == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func WriteErr(w http.ResponseWriter, err error) {
	var httpErr *httperr.HTTPError
	if !errors.As(err, &httpErr) {
		httpErr = httperr.InternalServerError(err.Error())
	}
	httpErr.Out(w)
}

func GetIntInPath(r *http.Request, name string) (int, error) {
	pathId := r.PathValue(name)
	id, err := strconv.Atoi(pathId)
	if err != nil {
		return 0, httperr.InputError("invalid id")
	}
	return id, nil
}

func BindJSON(r *http.Request, i any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return httperr.InputError("invalid content type")
	}
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		return httperr.InputError("invalid json body")
	}
	return nil
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
//	@Description		分页查询结果
type PaginationResult[T any] struct {
	Current int `json:"current"`
	Total   int `json:"total"`
	List    []T `json:"list"`
}

func NewPaginationResult[T any](current, total int, list []*T) *PaginationResult[T] {
	return &PaginationResult[T]{Current: current, Total: total, List: list}
}
