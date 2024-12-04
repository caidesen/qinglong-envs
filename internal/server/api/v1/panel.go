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
//
//	@Summary		创建面板
//	@Description	创建面板
//	@Tags			panel
//	@Accept			json
//	@Produce		json
//	@Param			input	body		queries.CreatePanelParams	true	"创建面板"
//	@Success		200
//	@Failure		400		{object}	httperr.HTTPError
//	@Failure		500		{object}	httperr.HTTPError
//	@Router			/api/panels [post]
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

// UpdatePanel 更新面板
//
//	@Summary		更新面板
//	@Description	更新面板
//	@Tags			panel
//	@Accept			json
//	@Produce		json
//	@Param			input	body		queries.UpdatePanelParams	true	"更新面板"
//	@Success		200		{object}	queries.Panel
//	@Failure		400		{object}	httperr.HTTPError
//	@Failure		500		{object}	httperr.HTTPError
//	@Router			/api/panels/{id} [put]
func (h *Handlers) UpdatePanel(w http.ResponseWriter, r *http.Request) {
	var input queries.UpdatePanelParams
	err := server.BindJSON(r, &input)
	if err != nil {
		server.WriteErr(w, err)
		return
	}
	err = h.queries.UpdatePanel(r.Context(), input)
	if err != nil {
		server.WriteErr(w, err)
		return
	}
}

// DeletePanel 删除面板
//
//	@Summary		删除面板
//	@Description	删除面板
//	@Tags			panel
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		400		{object}	httperr.HTTPError
//	@Failure		500		{object}	httperr.HTTPError
//	@Router			/api/panels/{id} [delete]
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

// GetPanelById 获取面板
//
//	@Summary		获取面板
//	@Description	获取面板
//	@Tags			panel
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	queries.Panel
//	@Failure		400		{object}	httperr.HTTPError
//	@Failure		500		{object}	httperr.HTTPError
//	@Router			/api/panels/{id} [get]
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

// ListPanels 获取面板列表
//
//	@Summary		获取面板列表
//	@Description	获取面板列表
//	@Tags			panel
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	server.PaginationResult[queries.Panel]
//	@Failure		400		{object}	httperr.HTTPError
//	@Failure		500		{object}	httperr.HTTPError
//	@Router			/api/panels [get]
func (h *Handlers) ListPanels(w http.ResponseWriter, r *http.Request) {
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
