package handler

import (
	"context"
	grpcUser "github.com/alfarezyyd/go-takemikazuchi-microservices-common/genproto/user"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/user/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userService service.Service
	grpcUser.UnimplementedUserServiceServer
}

func NewUserHandler(grpcServer *grpc.Server, userService service.Service) {
	categoryHandler := &UserHandler{
		userService: userService,
	}
	grpcUser.RegisterUserServiceServer(grpcServer, categoryHandler)
}

func (categoryHandler *UserHandler) CreateUser(ctx context.Context, categoryCreateRequest *grpcUser.CreateUserRequest) (*grpcUser.CommandUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (categoryHandler *UserHandler) UpdateUser(ctx context.Context, categoryCreateRequest *grpcUser.CreateUserRequest) (*grpcUser.CommandUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
