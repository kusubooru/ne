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

func (s *userService) GetAll(limit, offset int) ([]*ne.User, error) {
	u, err := s.Shimmie.GetAllUsers(limit, offset)
	if err != nil {
		return nil, err
	}
	users := make([]*ne.User, len(u))
	for i := range u {
		users[i] = &ne.User{
			ID:       u[i].ID,
			Name:     u[i].Name,
			Pass:     u[i].Pass,
			JoinDate: u[i].JoinDate,
			Admin:    u[i].Admin,
			Email:    u[i].Email,
			Class:    u[i].Class,
		}
	}
	return users, nil
}
