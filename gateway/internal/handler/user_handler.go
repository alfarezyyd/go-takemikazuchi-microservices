package handler

import (
	"context"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"time"
)

type UserHandler struct {
	grpcConnection *grpc.ClientConn
}

func NewUserHandler(grpcConnection *grpc.ClientConn) *UserHandler {
	return &UserHandler{
		grpcConnection: grpcConnection,
	}
}

func (userHandler *UserHandler) Register(ginContext *gin.Context) {
	userClient := user.NewUserServiceClient(userHandler.grpcConnection)
	var createUserDto userDto.CreateUserDto
	err := ginContext.ShouldBindBodyWithJSON(&createUserDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	createUserRequest := user.CreateUserRequest{
		Name:            createUserDto.Name,
		Email:           createUserDto.Email,
		PhoneNumber:     createUserDto.PhoneNumber,
		Password:        createUserDto.Password,
		ConfirmPassword: createUserDto.ConfirmPassword,
	}
	_, err = userClient.HandleRegister(timeoutCtx, &createUserRequest)
	fmt.Println(err)
	if err != nil {
		exception.ParseGrpcError(ginContext, err)
		return
	}
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("User created successfully", nil))
}

func (userHandler *UserHandler) GenerateOneTimePassword(ginContext *gin.Context) {
	var generateOneTimePassDto userDto.GenerateOtpDto
	err := ginContext.ShouldBindBodyWithJSON(&generateOneTimePassDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	//userHandler.userService.HandleGenerateOneTimePassword(&generateOneTimePassDto, nil)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("OTP generated successfully", nil))
}

func (userHandler *UserHandler) VerifyOneTimePassword(ginContext *gin.Context) {
	var VerifyOneTimePassDto userDto.VerifyOtpDto
	err := ginContext.ShouldBindBodyWithJSON(&VerifyOneTimePassDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	//userHandler.userService.HandleVerifyOneTimePassword(&VerifyOneTimePassDto)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("OTP verified successfully", nil))
}

func (userHandler *UserHandler) Login(ginContext *gin.Context) {
	userClient := user.NewUserServiceClient(userHandler.grpcConnection)
	var loginUserDto userDto.LoginUserDto
	err := ginContext.ShouldBindBodyWithJSON(&loginUserDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	payloadResponse, err := userClient.HandleLogin(timeoutCtx, &user.LoginUserRequest{
		UserIdentifier: loginUserDto.UserIdentifier,
		Password:       loginUserDto.Password,
	})
	if err != nil {
		exception.ParseGrpcError(ginContext, err)
		return
	}
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("User logged successfully", gin.H{
		"token": payloadResponse.Payload,
	}))
}

func (userHandler *UserHandler) LoginWithGoogle(ginContext *gin.Context) {
	userClient := user.NewUserServiceClient(userHandler.grpcConnection)
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	authEndpoint, _ := userClient.HandleGoogleAuthentication(timeout, &emptypb.Empty{})
	ginContext.Redirect(http.StatusSeeOther, authEndpoint.Payload)
	ginContext.JSON(http.StatusOK, authEndpoint)
}

func (userHandler *UserHandler) GoogleProviderCallback(ginContext *gin.Context) {
	userClient := user.NewUserServiceClient(userHandler.grpcConnection)
	tokenState := ginContext.Query("state")
	queryCode := ginContext.Query("code")
	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	_, err := userClient.HandleGoogleCallback(ctxTimeout, &user.GoogleCallbackRequest{
		TokenState: tokenState,
		QueryCode:  queryCode,
	})
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
}
