package rpc

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"

	"github.com/kusubooru/ne"
	"github.com/kusubooru/ne/jwt"
	"github.com/kusubooru/ne/jwt/csrf"
	"github.com/kusubooru/ne/rpc/pb"
	"github.com/kusubooru/ne/shimmie2"
)

type authServer struct {
	User   ne.UserService
	secret []byte
}

func NewAuthServer(userService ne.UserService, secret []byte) pb.AuthServer {
	return &authServer{User: userService, secret: secret}
}

func (s *authServer) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	if r == nil || r.PasswordHash == "" {
		return nil, fmt.Errorf("login credentials are required")
	}

	// Login User.
	_, err := s.User.Login(r.Username, r.PasswordHash)
	if err != nil {
		if err != shimmie2.ErrNotFound && err != shimmie2.ErrWrongCredentials {
			grpclog.Println("login failed:", err)
		}
		return nil, fmt.Errorf("login failed")
	}

	// Create CSRF token.
	csrfToken, err := csrf.NewToken()
	if err != nil {
		grpclog.Println("CSRF token creation failed:", err)
		return nil, fmt.Errorf("CSRF token creation failed")
	}
	grpc.SendHeader(ctx, metadata.Pairs("X-CSRF-Token", csrfToken))

	// Create Access token.
	userID := jwt.NewUUID()
	accessToken := &jwt.Token{
		Subject:  userID,
		Duration: time.Duration(15 * time.Minute),
		CSRF:     csrfToken,
	}
	signedAccessToken, err := jwt.Encode(accessToken, s.secret)
	if err != nil {
		grpclog.Println("access token creation failed:", err)
		return nil, fmt.Errorf("access token creation failed")
	}

	// Create Refresh token.
	refreshTokenID := jwt.NewUUID()
	refreshToken := &jwt.Token{
		ID:       refreshTokenID,
		Subject:  userID,
		Duration: time.Duration(72 * time.Hour),
		CSRF:     csrfToken,
	}
	signedRefreshToken, err := jwt.Encode(refreshToken, s.secret)
	if err != nil {
		grpclog.Println("refresh token creation failed:", err)
		return nil, fmt.Errorf("refresh token creation failed")
	}
	// TODO: Store refreshTokenID on cache/redis/db.

	resp := &pb.LoginResponse{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}
	return resp, nil
}

func (s *authServer) Refresh(ctx context.Context, r *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	return &pb.RefreshResponse{}, nil
}
