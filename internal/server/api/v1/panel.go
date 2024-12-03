package v1

import (
	"database/sql"
	"errors"
	"net/http"
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/internal/httperr"
	"qinglong-envs/internal/server"
)

// CreatePanel 创建面板
// @Summary 创建面板
// @Description 创建面板
// @Tags panel
// @Accept json
// @Produce json
// @Param input body queries.CreatePanelParams true "创建面板"
// @Success 200 {object} queries.Panel
// @Failure 400 {object} httperr.HTTPError
// @Failure 500 {object} httperr.HTTPError
// @Router /api/panels [post]
func (h *Handlers) CreatePanel(w http.ResponseWriter, r *http.Request) {
	var input queries.CreatePanelParams
	err := server.BindJSON(r, &input)
	if err != nil {
		server.WriteErr(w, err)
	}
	resp, err := h.queries.CreatePanel(r.Context(), input)
	if err != nil {
		server.WriteErr(w, err)
	}
	server.WriteJSON(w, resp)
}

func (h *Handlers) UpdatePanel(w http.ResponseWriter, r *http.Request) {
	var input queries.UpdatePanelParams
	err := server.BindJSON(r, &input)
	if err != nil {
		server.WriteErr(w, err)
	}
	err = h.queries.UpdatePanel(r.Context(), input)
	if err != nil {
		server.WriteErr(w, err)
	}
}

func (h *Handlers) DeletePanel(w http.ResponseWriter, r *http.Request) {
	id, err := server.GetIntInPath(r, "id")
	if err != nil {
		server.WriteErr(w, err)
	}
	err = h.queries.DeletePanelByID(r.Context(), id)
	if err != nil {
		server.WriteErr(w, err)
	}
}

func (h *Handlers) GetPanelById(w http.ResponseWriter, r *http.Request) {
	id, err := server.GetIntInPath(r, "id")
	if err != nil {
		server.WriteErr(w, err)
	}
	resp, err := h.queries.GetPanelByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httperr.NotFound("未找到面板").Out(w)
		}
		server.WriteErr(w, err)
	}
	server.WriteJSON(w, resp)
}

func (h *Handlers) ListPanels(w http.ResponseWriter, r *http.Request) {
	//current, _ := strconv.Atoi(r.URL.Query().Get("current"))
	//pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	//if current == -1 || pageSize < -1 || pageSize > 100 {
	//	handler.ValidateError("invalid pagination params").Out(w)
	//}
	//p := handler.PaginationParams{Current: current, PageSize: pageSize}
	panels, err := h.queries.ListPanels(
		r.Context(), queries.ListPanelsParams{
			Limit: 99, Offset: 0,
		})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			server.WriteJSON(w, server.NewPaginationResult(1, 1, []*any{}))
		}
		server.WriteErr(w, err)
	}
	server.WriteJSON(w, server.NewPaginationResult(1, 1, panels))
}
