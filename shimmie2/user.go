package shimmie2

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"reflect"

	"golang.org/x/net/context"

	"github.com/kusubooru/ne/ne"
	"github.com/kusubooru/shimmie"
)

const defaultLimit = 10

var (
	ErrWrongCredentials = errors.New("wrong username or password")
	ErrNotFound         = errors.New("entry not found")
)

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
	shimmieUsers, err := s.Shimmie.GetAllUsers(int(limit), int(offset))
	if err != nil {
		return nil, err
	}
	users := make([]*ne.User, len(shimmieUsers))
	for i := range shimmieUsers {
		users[i] = (*ne.User)(&shimmieUsers[i])
	}
	return users, nil
}

func (s *userService) Login(username, password string) (*ne.User, error) {
	u, err := s.Shimmie.GetUserByName(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	if u.Pass == shimmie.PasswordHash(username, password) {
		return (*ne.User)(u), nil
	}
	return nil, ErrWrongCredentials
}
