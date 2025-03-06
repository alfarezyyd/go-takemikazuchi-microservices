package handler

import (
	"context"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type UserHandler struct {
	userService       user.UserServiceClient
	validatorInstance *validator.Validate
}

func NewUserHandler(userService user.UserServiceClient, validatorInstance *validator.Validate) *UserHandler {
	return &UserHandler{
		userService:       userService,
		validatorInstance: validatorInstance,
	}
}

func (userHandler *UserHandler) Register(ginContext *gin.Context) {
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
	userHandler.userService.Register(timeoutCtx, &createUserRequest)
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
	var loginUserDto userDto.LoginUserDto
	err := ginContext.ShouldBindBodyWithJSON(&loginUserDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	_, err = userHandler.userService.HandleLogin(timeoutCtx, &user.LoginUserRequest{
		UserIdentifier: loginUserDto.UserIdentifier,
		Password:       loginUserDto.Password,
	})
	fmt.Println(err)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("User logged successfully", gin.H{
		"token": "ON TESTING",
	}))
}
