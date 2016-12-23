package ne

import "github.com/kusubooru/shimmie"

type UserService interface {
	GetAll(limit, offset int) ([]shimmie.User, error)
}
