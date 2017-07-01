package rpc

import (
	"golang.org/x/net/context"

	"github.com/kusubooru/ne/ne"
	"github.com/kusubooru/ne/rpc/pb"
)

const defaultPageSize = 10

type usersServer struct {
	User ne.UserService
}

func NewUsersServer(userService ne.UserService) pb.UsersServer {
	return &usersServer{User: userService}
}

func newUser(u *ne.User) *pb.User {
	return &pb.User{
		Id:   u.ID,
		Name: u.Name,
	}
}

func (s *usersServer) StreamAll(page *pb.Page, stream pb.Users_StreamAllServer) error {
	limit, offset := toOffset(page)

	users, err := s.User.GetAll(stream.Context(), limit, offset)
	if err != nil {
		return err
	}
	for _, u := range users {
		pbu := newUser(u)
		if err := stream.Send(pbu); err != nil {
			return err
		}
	}
	return nil
}

func (s *usersServer) GetAll(ctx context.Context, page *pb.Page) (*pb.GetAllResponse, error) {
	limit, offset := toOffset(page)

	users, err := s.User.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	pbUsers := make([]*pb.User, len(users))
	for i := range users {
		pbUsers[i] = newUser(users[i])
	}
	return &pb.GetAllResponse{Users: pbUsers}, nil
}

func toOffset(page *pb.Page) (limit, offset int64) {
	if page == nil {
		return 0, 0
	}
	limit, offset = page.Limit, page.Offset
	if page.Page != 0 {
		limit, offset = pageToOffset(page.Page, page.PerPage)
		return
	}
	if page.PerPage != 0 {
		limit, offset = pageToOffset(page.Page, page.PerPage)
		return
	}
	return
}

func pageToOffset(page, perPage int64) (limit, offset int64) {
	if page < 1 {
		page = 1
	}
	if perPage == 0 {
		perPage = defaultPageSize
	}
	offset = (page - 1) * perPage
	limit = perPage
	return limit, offset
}
