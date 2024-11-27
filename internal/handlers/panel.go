package handlers

import (
	"net/http"
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/internal/services"
	"qinglong-envs/pkg/api"
	"qinglong-envs/pkg/router"
	"strconv"
)

type PanelHandler struct {
	panelServer *services.PanelServer
}

func NewPanelHandler(panelServer *services.PanelServer) *PanelHandler {
	return &PanelHandler{
		panelServer: panelServer,
	}
}

func (h *PanelHandler) Register(r router.Router) {
	r.Post("/panels", api.Handler(h.createPanel))
	r.Put("/panels", api.Handler(h.updatePanel))
	r.Delete("/panels/{id}", api.Handler(h.deletePanel))
	r.Get("/panels/{id}", api.Handler(h.getPanel))
	r.Get("/panels", api.Handler(h.listPanels))
}

func (h *PanelHandler) createPanel(w http.ResponseWriter, r *http.Request) error {
	input, err := api.BindJSONBody[queries.CreatePanelParams](r)
	if err != nil {
		return err
	}
	resp, err := h.panelServer.CreatePanel(r.Context(), input)
	if err != nil {
		return err
	}
	return api.Response(w, resp)
}

func (h *PanelHandler) updatePanel(_ http.ResponseWriter, r *http.Request) error {
	input, err := api.BindJSONBody[queries.UpdatePanelParams](r)
	if err != nil {
		return err
	}
	return h.panelServer.UpdatePanel(r.Context(), input)
}

func (h *PanelHandler) deletePanel(_ http.ResponseWriter, r *http.Request) error {
	id, err := api.GetIntInPath(r, "id")
	if err != nil {
		return err
	}
	return h.panelServer.DeletePanel(r.Context(), id)
}

func (h *PanelHandler) getPanel(w http.ResponseWriter, r *http.Request) error {
	id, err := api.GetIntInPath(r, "id")
	if err != nil {
		return err
	}
	resp, err := h.panelServer.GetPanelById(r.Context(), id)
	if err != nil {
		return err
	}
	return api.Response(w, resp)
}

func (h *PanelHandler) listPanels(w http.ResponseWriter, r *http.Request) error {
	current, _ := strconv.Atoi(r.URL.Query().Get("current"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if current == -1 || pageSize < -1 || pageSize > 100 {
		return api.ValidateError("invalid pagination params")
	}
	resp, err := h.panelServer.ListPanels(r.Context(), api.PaginationParams{Current: current, PageSize: pageSize})
	if err != nil {
		return err
	}
	return api.Response(w, resp)
}
