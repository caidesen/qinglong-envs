package apiv1

import (
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/pkg/httpapi"
)

type apiv1 struct {
	queries *queries.Queries
}

func New(queries *queries.Queries) *apiv1 {
	return &apiv1{queries: queries}
}

func (a *apiv1) Register(r httpapi.Router) {
	r.Group(func(router httpapi.Router) {
		r.API("GET /panels/{id}", a.GetPanelById)
		r.API("GET /panels", a.ListPanels)
		r.API("POST /panels", a.CreatePanel)
		r.API("PUT /panels", a.UpdatePanel)
		r.API("DELETE /panels/{id}", a.DeletePanel)
	})
}
