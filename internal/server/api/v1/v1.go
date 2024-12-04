package apiv1

import (
	"github.com/go-playground/validator/v10"
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/internal/server"
	"qinglong-envs/pkg/router"
)

type Handlers struct {
	queries *queries.Queries
	*server.Handler
}

func NewHandlers(queries *queries.Queries, validate *validator.Validate) *Handlers {
	return &Handlers{queries: queries, Handler: server.NewHandler(validate)}
}

func (h *Handlers) Routes(r router.Router) {
	{
		r.HandleFunc("GET /panels/{id}", h.GetPanelById)
		r.HandleFunc("GET /panels", h.ListPanels)
		r.HandleFunc("POST /panels", h.CreatePanel)
		r.HandleFunc("PUT /panels", h.UpdatePanel)
		r.HandleFunc("DELETE /panels/{id}", h.DeletePanel)
	}
}
