package apiv1

import (
	"database/sql"
	"errors"
	"net/http"
	"qinglong-envs/internal/db/queries"
	"qinglong-envs/internal/httperr"
)

type (
	CreatePanelInput struct {
		Name         string `json:"name" validate:"required,min=1,max=64"`
		Url          string `json:"url" validate:"required,url"`
		ClientID     string `json:"clientId" validate:"required,min=1,max=128"`
		ClientSecret string `json:"clientSecret" validate:"required,min=1,max=128"`
	}

	UpdatePanelInput struct {
		ID int `json:"id" validate:"required"`
		CreatePanelInput
	}
)

func (i *CreatePanelInput) Q() queries.CreatePanelParams {
	return queries.CreatePanelParams{
		Name:         i.Name,
		Url:          i.Url,
		ClientID:     i.ClientID,
		ClientSecret: i.ClientSecret,
	}
}
func (i *UpdatePanelInput) Q() queries.UpdatePanelParams {
	return queries.UpdatePanelParams{
		ID:           i.ID,
		Name:         i.Name,
		Url:          i.Url,
		ClientID:     i.ClientID,
		ClientSecret: i.ClientSecret,
	}
}

// CreatePanel 创建面板
//
//	@Summary		创建面板
//	@Description	创建面板
//	@Tags			panel
//	@Accept			json
//	@Produce		json
//	@Param			input	body	CreatePanelInput	true	"创建面板"
//	@Success		200
//	@Failure		400	{object}	httperr.HTTPError
//	@Failure		500	{object}	httperr.HTTPError
//	@Router			/panels [post]
func (h *Handlers) CreatePanel(w http.ResponseWriter, r *http.Request) {
	var input CreatePanelInput
	err := h.BindJSON(r, &input)
	if err != nil {
		h.WriteErr(w, err)
		return
	}
	err = h.Validate(input)
	if err != nil {
		h.WriteErr(w, err)
		return
	}
	resp, err := h.queries.CreatePanel(r.Context(), input.Q())
	if err != nil {
		h.WriteErr(w, err)
		return
	}
	h.WriteJSON(w, resp)
}

// UpdatePanel 更新面板
//
//	@Summary		更新面板
//	@Description	更新面板
//	@Tags			panel
//	@Accept			json
//	@Produce		json
//	@Param			input	body		UpdatePanelInput	true	"创建面板"
//	@Success		200		{object}	queries.Panel
//	@Failure		400		{object}	httperr.HTTPError
//	@Failure		500		{object}	httperr.HTTPError
//	@Router			/panels [put]
func (h *Handlers) UpdatePanel(w http.ResponseWriter, r *http.Request) {
	var input UpdatePanelInput
	err := h.BindJSON(r, &input)
	if err != nil {
		h.WriteErr(w, err)
		return
	}
	err = h.queries.UpdatePanel(r.Context(), input.Q())
	if err != nil {
		h.WriteErr(w, err)
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
//	@Param			id	path	int	true	"面板ID"
//	@Success		200
//	@Failure		400	{object}	httperr.HTTPError
//	@Failure		500	{object}	httperr.HTTPError
//	@Router			/panels/{id} [delete]
func (h *Handlers) DeletePanel(w http.ResponseWriter, r *http.Request) {
	id, err := h.GetIntInPath(r, "id")
	if err != nil {
		h.WriteErr(w, err)
		return
	}
	err = h.queries.DeletePanelByID(r.Context(), id)
	if err != nil {
		h.WriteErr(w, err)
		return
	}
}

// GetPanelById 获取面板
//
//	@Summary		获取面板
//	@Description	获取面板
//	@Tags			panel
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"面板ID"
//	@Success		200	{object}	queries.Panel
//	@Failure		400	{object}	httperr.HTTPError
//	@Failure		500	{object}	httperr.HTTPError
//	@Router			/panels/{id} [get]
func (h *Handlers) GetPanelById(w http.ResponseWriter, r *http.Request) {
	id, err := h.GetIntInPath(r, "id")
	if err != nil {
		h.WriteErr(w, err)
		return
	}
	resp, err := h.queries.GetPanelByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httperr.NotFound("未找到面板").Out(w)
			return
		}
		h.WriteErr(w, err)
		return
	}
	h.WriteJSON(w, resp)
	return
}

// ListPanels 获取面板列表
//
//	@Summary		获取面板列表
//	@Description	获取面板列表
//	@Tags			panel
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		queries.Panel
//	@Failure		400	{object}	httperr.HTTPError
//	@Failure		500	{object}	httperr.HTTPError
//	@Router			/panels [get]
func (h *Handlers) ListPanels(w http.ResponseWriter, r *http.Request) {
	panels, err := h.queries.ListPanels(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.WriteJSON(w, []*any{})
			return
		}
		h.WriteErr(w, err)
		return
	}
	h.WriteJSON(w, panels)
}
