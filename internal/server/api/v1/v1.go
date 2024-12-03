package v1

import (
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/pkg/router"
)

type Handlers struct {
	queries *queries.Queries
}

func NewHandlers(queries *queries.Queries) *Handlers {
	return &Handlers{queries: queries}
}

func (h *Handlers) Routes(r router.Router) {
	api := r.Mount("/api/v1")
	api.Route(func(r router.Router) {
		r.HandleFunc("GET /panels/{id}", h.GetPanelById)
		r.HandleFunc("GET /panels", h.ListPanels)
		r.HandleFunc("POST /panels", h.CreatePanel)
		r.HandleFunc("PUT /panels", h.UpdatePanel)
		r.HandleFunc("DELETE /panels/{id}", h.DeletePanel)
	})
}
