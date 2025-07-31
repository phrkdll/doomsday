package auth

import (
	"context"
	"errors"
	"log/slog"

	pb "github.com/phrkdll/doomsday/protos"
)

var (
	ErrInvalidRequest = errors.New("invalid request: username and password are required")
	ErrNotImplemented = errors.New("not implemented")
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthService) Register(c context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if request == nil || request.Username == "" || request.Password == "" {
		return nil, ErrInvalidRequest
	}

	slog.Info("Register called", "username", request.Username, "password", request.Password)

	return &pb.RegisterResponse{Message: "success", Token: "Yes"}, nil
}

func (s *AuthService) Login(c context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, ErrNotImplemented
}

func (s *AuthService) Logout(c context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, ErrNotImplemented
}
