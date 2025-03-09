package handler

import (
	"context"
	grpcUserAddress "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user_address"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/service"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserAddressHandler struct {
	userAddressService service.UserAddressService
	grpcUserAddress.UnimplementedUserAddressServiceServer
}

func NewUserAddressHandler(grpcServer *grpc.Server, userAddressService service.UserAddressService) {
	userAddressHandler := &UserAddressHandler{
		userAddressService: userAddressService,
	}
	grpcUserAddress.RegisterUserAddressServiceServer(grpcServer, userAddressHandler)
}

func (userAddressHandler *UserAddressHandler) UserAddressStore(ctx context.Context, userAddressCreateRequest *grpcUserAddress.UserAddressCreateRequest) (*grpcUserAddress.QueryResponse, error) {
	userAddressId := userAddressHandler.userAddressService.Create(ctx, &dto.CreateUserAddressDto{
		Latitude:  userAddressCreateRequest.Latitude,
		Longitude: userAddressCreateRequest.Longitude,
		UserId:    userAddressCreateRequest.UserId,
	})
	return &grpcUserAddress.QueryResponse{Id: userAddressId}, nil
}
func (userAddressHandler *UserAddressHandler) FindUserAddressById(ctx context.Context, userAddressSearchRequest *grpcUserAddress.UserAddressSearchRequest) (*grpcUserAddress.QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserAddressById not implemented")
}
