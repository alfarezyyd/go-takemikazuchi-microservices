package user

import (
	"go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
)

type Service interface {
	HandleRegister(createUserDto *dto.CreateUserDto)
	HandleGenerateOneTimePassword(generateOtpDto *dto.GenerateOtpDto)
	HandleVerifyOneTimePassword(verifyOtpDto *dto.VerifyOtpDto)
	HandleGoogleAuthentication() string
	HandleGoogleCallback(tokenState string, queryCode string) *exception.ClientError
	HandleLogin(loginUserDto *dto.LoginUserDto) string
}
