package handler

import (
	"context"
	grpcUser "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/user/service"
	"google.golang.org/grpc"
)

type UserHandler struct {
	userService service.Service
	grpcUser.UnimplementedUserServiceServer
}

func NewUserHandler(grpcServer *grpc.Server, userService service.Service) {
	userHandler := &UserHandler{
		userService: userService,
	}
	grpcUser.RegisterUserServiceServer(grpcServer, userHandler)
}

func (h *UserHandler) HandleLogin(ctx context.Context, req *grpcUser.LoginUserRequest) (*grpcUser.LoginResponse, error) {
	return h.userService.HandleLogin(ctx, req)
}
