package user

import (
	dto2 "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
)

type Service interface {
	HandleRegister(createUserDto *dto2.CreateUserDto)
	HandleGenerateOneTimePassword(generateOtpDto *dto2.GenerateOtpDto)
	HandleVerifyOneTimePassword(verifyOtpDto *dto2.VerifyOtpDto)
	HandleGoogleAuthentication() string
	HandleGoogleCallback(tokenState string, queryCode string) *exception.ClientError
	HandleLogin(loginUserDto *dto2.LoginUserDto) string
}
