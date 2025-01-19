package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto2 "go-takemikazuchi-api/internal/user/dto"
	exception2 "go-takemikazuchi-api/pkg/exception"
	helper2 "go-takemikazuchi-api/pkg/helper"
	"net/http"
)

type Handler struct {
	userService       Service
	validatorInstance *validator.Validate
}

func NewHandler(userService Service, validatorInstance *validator.Validate) *Handler {
	return &Handler{
		userService:       userService,
		validatorInstance: validatorInstance,
	}
}

func (userHandler *Handler) Register(ginContext *gin.Context) {
	var createUserDto dto2.CreateUserDto
	err := ginContext.ShouldBindBodyWithJSON(&createUserDto)
	helper2.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
	userHandler.userService.HandleRegister(&createUserDto)
	ginContext.JSON(http.StatusOK, helper2.WriteSuccess("User created successfully", nil))
}

func (userHandler *Handler) GenerateOneTimePassword(ginContext *gin.Context) {
	var generateOneTimePassDto dto2.GenerateOtpDto
	err := ginContext.ShouldBindBodyWithJSON(&generateOneTimePassDto)
	helper2.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
	userHandler.userService.HandleGenerateOneTimePassword(&generateOneTimePassDto)
	ginContext.JSON(http.StatusOK, helper2.WriteSuccess("OTP generated successfully", nil))
}

func (userHandler *Handler) VerifyOneTimePassword(ginContext *gin.Context) {
	var VerifyOneTimePassDto dto2.VerifyOtpDto
	err := ginContext.ShouldBindBodyWithJSON(&VerifyOneTimePassDto)
	helper2.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
	userHandler.userService.HandleVerifyOneTimePassword(&VerifyOneTimePassDto)
	ginContext.JSON(http.StatusOK, helper2.WriteSuccess("OTP verified successfully", nil))
}

func (userHandler *Handler) Login(ginContext *gin.Context) {
	var loginUserDto dto2.LoginUserDto
	err := ginContext.ShouldBindBodyWithJSON(&loginUserDto)
	helper2.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
	generatedToken := userHandler.userService.HandleLogin(&loginUserDto)
	ginContext.JSON(http.StatusOK, helper2.WriteSuccess("User logged successfully", gin.H{
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
	helper2.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
}
