package rest

import (
	"net/http"

	"github.com/kusubooru/ne/ne"
)

const contentType = "application/vnd.api+json"

// API holds handlers for each resource.
type API struct {
	User *UserHandler
}

// New creates a new API.
func New(userService ne.UserService) *API {
	api := &API{
		User: NewUserHandler(userService),
	}
	return api
}

func (api *API) Handlers() *http.ServeMux {
	mux := http.NewServeMux()

	api.User.Register(mux)

	return mux
}
