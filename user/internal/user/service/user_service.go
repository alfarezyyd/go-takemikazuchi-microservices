package service

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/user/dto"
	"gorm.io/gorm"
)

type Service interface {
	HandleRegister(createUserDto *dto.CreateUserDto)
	HandleGenerateOneTimePassword(generateOtpDto *dto.GenerateOtpDto, externalGormTransaction *gorm.DB)
	HandleVerifyOneTimePassword(verifyOtpDto *dto.VerifyOtpDto)
	HandleGoogleAuthentication() string
	HandleGoogleCallback(tokenState string, queryCode string) *exception.ClientError
	HandleLogin(loginUserDto *dto.LoginUserDto) string
}
