package rest_test

import (
	"reflect"
	"testing"

	"github.com/kusubooru/ne"
	"github.com/kusubooru/ne/rest"
)

var users = []*ne.User{
	{ID: 1, Name: "john"},
	{ID: 2, Name: "mary"},
}

type userService struct{}

func (s *userService) GetAll(limit, offset int) ([]*ne.User, error) {
	return users, nil
}

func TestNew(t *testing.T) {
	s := &userService{}
	h := &rest.UserHandler{Service: s}
	want := &rest.API{User: h}
	got := rest.New(s)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("rest.New(%#v) = \n%#v, want \n%#v", s, got, want)
	}
}
