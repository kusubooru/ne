package shimmie2

import (
	"fmt"
	"math"
	"reflect"

	"golang.org/x/net/context"

	"github.com/kusubooru/ne"
	"github.com/kusubooru/shimmie"
)

const defaultLimit = 10

type userService struct {
	Shimmie shimmie.Store
}

func NewUserService(s shimmie.Store) ne.UserService {
	return &userService{Shimmie: s}
}

func (s *userService) GetAll(ctx context.Context, limit, offset int64) ([]*ne.User, error) {
	if limit == 0 {
		limit = defaultLimit
	}
	var a int
	// Temporarily checking for int size till shimmie function changes to
	// int64. int must be 8 bytes to be able to handle int64.
	if reflect.TypeOf(a).Size() < 8 {
		if limit > math.MaxInt32 || offset > math.MaxInt32 {
			return nil, fmt.Errorf("limit or offset too large for current platform")
		}
	}
	u, err := s.Shimmie.GetAllUsers(int(limit), int(offset))
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
