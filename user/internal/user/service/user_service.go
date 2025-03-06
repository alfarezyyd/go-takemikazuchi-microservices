package service

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"gorm.io/gorm"
)

type Service interface {
	HandleRegister(createUserDto *userDto.CreateUserDto)
	HandleGenerateOneTimePassword(generateOtpDto *userDto.GenerateOtpDto, externalGormTransaction *gorm.DB)
	HandleVerifyOneTimePassword(verifyOtpDto *userDto.VerifyOtpDto)
	HandleGoogleAuthentication() string
	HandleGoogleCallback(tokenState string, queryCode string) *exception.ClientError
	HandleLogin(loginUserDto *userDto.LoginUserDto) string
}
