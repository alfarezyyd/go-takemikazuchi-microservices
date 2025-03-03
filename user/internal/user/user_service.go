package user

import (
	"go-takemikazuchi-microservices/internal/user/dto"
	"go-takemikazuchi-microservices/pkg/exception"
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
