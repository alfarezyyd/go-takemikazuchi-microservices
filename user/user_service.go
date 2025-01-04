package user

import (
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/user/dto"
)

type Service interface {
	HandleRegister(createUserDto *dto.CreateUserDto)
	HandleGenerateOneTimePassword(generateOtpDto *dto.GenerateOtpDto)
	HandleVerifyOneTimePassword(verifyOtpDto *dto.VerifyOtpDto)
	HandleGoogleAuthentication() string
	HandleGoogleCallback(tokenState string, queryCode string) *exception.ClientError
	HandleLogin(loginUserDto *dto.LoginUserDto) string
}
