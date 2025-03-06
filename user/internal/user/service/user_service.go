package service

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/exception"
	dto2 "github.com/alfarezyyd/go-takemikazuchi-microservices-user/pkg/dto"
	"gorm.io/gorm"
)

type Service interface {
	HandleRegister(createUserDto *dto2.CreateUserDto)
	HandleGenerateOneTimePassword(generateOtpDto *dto2.GenerateOtpDto, externalGormTransaction *gorm.DB)
	HandleVerifyOneTimePassword(verifyOtpDto *dto2.VerifyOtpDto)
	HandleGoogleAuthentication() string
	HandleGoogleCallback(tokenState string, queryCode string) *exception.ClientError
	HandleLogin(loginUserDto *dto2.LoginUserDto) string
}
