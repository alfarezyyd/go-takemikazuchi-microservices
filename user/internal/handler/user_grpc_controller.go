package handler

import (
	"context"
	grpcUser "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/service"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	userService service.UserService
	grpcUser.UnimplementedUserServiceServer
}

func NewUserHandler(grpcServer *grpc.Server, userService service.UserService) {
	userHandler := &UserHandler{
		userService: userService,
	}
	grpcUser.RegisterUserServiceServer(grpcServer, userHandler)
}

func (userHandler *UserHandler) HandleLogin(ctx context.Context, req *grpcUser.LoginUserRequest) (*grpcUser.PayloadResponse, error) {
	tokenString := userHandler.userService.HandleLogin(req)
	return &grpcUser.PayloadResponse{
		Payload: tokenString,
	}, nil
}

func (userHandler *UserHandler) HandleRegister(ctx context.Context, req *grpcUser.CreateUserRequest) (*grpcUser.CommandUserResponse, error) {
	err := userHandler.userService.HandleRegister(&dto.CreateUserDto{
		Name:            req.Name,
		Email:           req.Email,
		PhoneNumber:     req.PhoneNumber,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		return nil, err
	}
	return &grpcUser.CommandUserResponse{
		IsSuccess: true,
	}, nil
}

func (userHandler *UserHandler) HandleVerifyOneTimePassword(ctx context.Context, req *grpcUser.VerifyOtpRequest) (*grpcUser.QueryUserResponse, error) {
	userHandler.userService.HandleVerifyOneTimePassword(&dto.VerifyOtpDto{
		Email:                req.Email,
		OneTimePasswordToken: req.OneTimePasswordToken,
	})
	return nil, nil
}

func (userHandler *UserHandler) HandleGoogleAuthentication(ctx context.Context, emptyProto *emptypb.Empty) (*grpcUser.PayloadResponse, error) {
	authenticationString := userHandler.userService.HandleGoogleAuthentication()
	return &grpcUser.PayloadResponse{
		Payload: authenticationString,
	}, nil
}

func (userHandler *UserHandler) HandleGoogleCallback(ctx context.Context, googleCallbackRequest *grpcUser.GoogleCallbackRequest) (*emptypb.Empty, error) {
	err := userHandler.userService.HandleGoogleCallback(googleCallbackRequest.TokenState, googleCallbackRequest.QueryCode)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (userHandler *UserHandler) FindByIdentifier(ctx context.Context, userIdentifier *grpcUser.UserIdentifier) (*grpcUser.QueryUserResponse, error) {
	userResponseDto := userHandler.userService.FindByIdentifier(&dto.UserIdentifierDto{
		Email:       userIdentifier.Email,
		PhoneNumber: userIdentifier.PhoneNumber,
	})
	return mapper.MapUserResponseIntoQueryUserResponse(userResponseDto), nil
}
