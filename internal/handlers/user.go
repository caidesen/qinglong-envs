package handlers

import (
	"net/http"
	"qinglong-envs/internal/services"
	"qinglong-envs/pkg/api"
	"qinglong-envs/pkg/router"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(r router.Router) {
	r.Post("/users/login", api.Handler(h.loginByLocal))
	r.Post("/users/register", api.Handler(h.register))
}

func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) error {
	input, err := api.BindJSONBody[services.UsernamePasswordInput](r)
	if err != nil {
		return err
	}
	resp, err := h.userService.Register(r.Context(), input)
	if err != nil {
		return err
	}
	return api.JSON(w, resp)
}

func (h *UserHandler) loginByLocal(w http.ResponseWriter, r *http.Request) error {
	input, err := api.BindJSONBody[services.UsernamePasswordInput](r)
	if err != nil {
		return err
	}
	resp, err := h.userService.LoginByLocal(r.Context(), input)
	if err != nil {
		return err
	}
	return api.JSON(w, resp)
}
