package apiv1

import (
	"database/sql"
	"errors"
	"net/http"
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/pkg/httpapi"
	"strconv"
)

//type PanelHandler struct {
//	panelServer *services.PanelServer
//}
//
//func NewPanelHandler(panelServer *services.PanelServer) *PanelHandler {
//	return &PanelHandler{
//		panelServer: panelServer,
//	}
//}

//func (h *PanelHandler) R(r httpapi.Router) {
//	r.Post("/panels", httpapi.HandlerFunc(h.Create))
//	r.Put("/panels", httpapi.HandlerFunc(h.Update))
//	r.Delete("/panels/{id}", httpapi.HandlerFunc(h.Delete))
//	r.Get("/panels/{id}", httpapi.HandlerFunc(h.GetOne))
//	r.Get("/panels", httpapi.HandlerFunc(h.List))
//}

func (a *apiv1) CreatePanel(w http.ResponseWriter, r *http.Request) error {
	input, err := httpapi.BindJSONBody[queries.CreatePanelParams](r)
	if err != nil {
		return err
	}
	resp, err := a.queries.CreatePanel(r.Context(), input)
	if err != nil {
		return err
	}
	return httpapi.Response(w, resp)
}

func (a *apiv1) UpdatePanel(_ http.ResponseWriter, r *http.Request) error {
	input, err := httpapi.BindJSONBody[queries.UpdatePanelParams](r)
	if err != nil {
		return err
	}
	return a.queries.UpdatePanel(r.Context(), input)
}

func (a *apiv1) DeletePanel(_ http.ResponseWriter, r *http.Request) error {
	id, err := httpapi.GetIntInPath(r, "id")
	if err != nil {
		return err
	}
	return a.queries.DeletePanelByID(r.Context(), id)
}

func (a *apiv1) GetPanelById(w http.ResponseWriter, r *http.Request) error {
	id, err := httpapi.GetIntInPath(r, "id")
	if err != nil {
		return err
	}
	resp, err := a.queries.GetPanelByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return httpapi.NotFoundError("未找到面板")
		}
		return err
	}
	return httpapi.Response(w, resp)
}

func (a *apiv1) ListPanels(w http.ResponseWriter, r *http.Request) error {
	current, _ := strconv.Atoi(r.URL.Query().Get("current"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if current == -1 || pageSize < -1 || pageSize > 100 {
		return httpapi.ValidateError("invalid pagination params")
	}
	p := httpapi.PaginationParams{Current: current, PageSize: pageSize}
	panels, err := a.queries.ListPanels(
		r.Context(), queries.ListPanelsParams{
			Limit: p.Current, Offset: p.Offset(),
		})
	if err != nil {
		return err
	}
	return httpapi.Response(w, httpapi.NewPaginationResult(1, 1, panels))
}
