package rest

import (
	"net/http"

	"github.com/kusubooru/ne"
)

const contentType = "application/vnd.api+json"

type API struct {
	User *UserHandler
}

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
