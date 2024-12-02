package apiv1

import (
	"net/http"
	"qinglong-envs/internal/services"
	"qinglong-envs/pkg/httpapi"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

//func (h *UserHandler) R(r httpapi.Router) {
//	r.Post("/users/login", httpapi.HandlerFunc(h.loginByLocal))
//	r.Post("/users/register", httpapi.HandlerFunc(h.register))
//}

func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) error {
	input, err := httpapi.BindJSONBody[services.UsernamePasswordInput](r)
	if err != nil {
		return err
	}
	resp, err := h.userService.Register(r.Context(), input)
	if err != nil {
		return err
	}
	return httpapi.Response(w, resp)
}

func (h *UserHandler) loginByLocal(w http.ResponseWriter, r *http.Request) error {
	input, err := httpapi.BindJSONBody[services.UsernamePasswordInput](r)
	if err != nil {
		return err
	}
	resp, err := h.userService.LoginByLocal(r.Context(), input)
	if err != nil {
		return err
	}
	return httpapi.Response(w, resp)
}
