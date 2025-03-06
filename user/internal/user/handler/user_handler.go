package handler

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/user/service"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-user/pkg/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	userService       service.Service
	validatorInstance *validator.Validate
}

func NewHandler(userService service.Service, validatorInstance *validator.Validate) *Handler {
	return &Handler{
		userService:       userService,
		validatorInstance: validatorInstance,
	}
}

func (userHandler *Handler) Register(ginContext *gin.Context) {
	var createUserDto dto.CreateUserDto
	err := ginContext.ShouldBindBodyWithJSON(&createUserDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	userHandler.userService.HandleRegister(&createUserDto)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("User created successfully", nil))
}

func (userHandler *Handler) GenerateOneTimePassword(ginContext *gin.Context) {
	var generateOneTimePassDto dto.GenerateOtpDto
	err := ginContext.ShouldBindBodyWithJSON(&generateOneTimePassDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	userHandler.userService.HandleGenerateOneTimePassword(&generateOneTimePassDto, nil)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("OTP generated successfully", nil))
}

func (userHandler *Handler) VerifyOneTimePassword(ginContext *gin.Context) {
	var VerifyOneTimePassDto dto.VerifyOtpDto
	err := ginContext.ShouldBindBodyWithJSON(&VerifyOneTimePassDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	userHandler.userService.HandleVerifyOneTimePassword(&VerifyOneTimePassDto)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("OTP verified successfully", nil))
}

func (userHandler *Handler) Login(ginContext *gin.Context) {
	var loginUserDto dto.LoginUserDto
	err := ginContext.ShouldBindBodyWithJSON(&loginUserDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	generatedToken := userHandler.userService.HandleLogin(&loginUserDto)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("User logged successfully", gin.H{
		"token": generatedToken,
	}))
}

func (userHandler *Handler) LoginWithGoogle(ginContext *gin.Context) {
	authEndpoint := userHandler.userService.HandleGoogleAuthentication()
	ginContext.Redirect(http.StatusSeeOther, authEndpoint)
	ginContext.JSON(http.StatusOK, authEndpoint)
}

func (userHandler *Handler) GoogleProviderCallback(ginContext *gin.Context) {
	tokenState := ginContext.Query("state")
	queryCode := ginContext.Query("code")
	err := userHandler.userService.HandleGoogleCallback(tokenState, queryCode)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
}
