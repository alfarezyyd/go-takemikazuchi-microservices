package handler

import (
	"context"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"time"
)

type UserHandler struct {
	serviceRegistry discovery.ServiceRegistry
}

func NewUserHandler(serviceRegistry discovery.ServiceRegistry,
) *UserHandler {
	return &UserHandler{
		serviceRegistry: serviceRegistry,
	}
}

func (userHandler *UserHandler) Register(ginContext *gin.Context) {
	var createUserDto dto.CreateUserDto
	err := ginContext.ShouldBindBodyWithJSON(&createUserDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "userService", userHandler.serviceRegistry)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	userClient := user.NewUserServiceClient(grpcConnection)
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
	var generateOneTimePassDto dto.GenerateOtpDto
	err := ginContext.ShouldBindBodyWithJSON(&generateOneTimePassDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	//userHandler.userService.HandleGenerateOneTimePassword(&generateOneTimePassDto, nil)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("OTP generated successfully", nil))
}

func (userHandler *UserHandler) VerifyOneTimePassword(ginContext *gin.Context) {
	var VerifyOneTimePassDto dto.VerifyOtpDto
	err := ginContext.ShouldBindBodyWithJSON(&VerifyOneTimePassDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	//userHandler.userService.HandleVerifyOneTimePassword(&VerifyOneTimePassDto)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("OTP verified successfully", nil))
}

func (userHandler *UserHandler) Login(ginContext *gin.Context) {
	var loginUserDto dto.LoginUserDto
	err := ginContext.ShouldBindBodyWithJSON(&loginUserDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "userService", userHandler.serviceRegistry)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	userClient := user.NewUserServiceClient(grpcConnection)

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
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeout, "userService", userHandler.serviceRegistry)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	userClient := user.NewUserServiceClient(grpcConnection)
	authEndpoint, _ := userClient.HandleGoogleAuthentication(timeout, &emptypb.Empty{})
	ginContext.Redirect(http.StatusSeeOther, authEndpoint.Payload)
	ginContext.JSON(http.StatusOK, authEndpoint)
}

func (userHandler *UserHandler) GoogleProviderCallback(ginContext *gin.Context) {
	tokenState := ginContext.Query("state")
	queryCode := ginContext.Query("code")
	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(ctxTimeout, "userService", userHandler.serviceRegistry)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	userClient := user.NewUserServiceClient(grpcConnection)
	_, err = userClient.HandleGoogleCallback(ctxTimeout, &user.GoogleCallbackRequest{
		TokenState: tokenState,
		QueryCode:  queryCode,
	})
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
}
