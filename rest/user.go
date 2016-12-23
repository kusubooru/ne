package rest

import (
	"encoding/json"
	"net/http"

	"github.com/kusubooru/ne"
)

type UserHandler struct {
	Service ne.UserService
}

func NewUserHandler(s ne.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) Register(mux *http.ServeMux) {
	mux.Handle("/users", h)
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAll(10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
