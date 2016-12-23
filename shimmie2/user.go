package shimmie2

import (
	"github.com/kusubooru/ne"
	"github.com/kusubooru/shimmie"
)

type userService struct {
	Shimmie shimmie.Store
}

func NewUserService(s shimmie.Store) ne.UserService {
	return &userService{Shimmie: s}
}

func (s *userService) GetAll(limit, offset int) ([]shimmie.User, error) {
	return s.Shimmie.GetAllUsers(limit, offset)
}
