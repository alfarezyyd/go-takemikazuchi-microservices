package service

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"gorm.io/gorm"
)

type Service interface {
	HandleRegister(context.Context, *user.CreateUserRequest) (*user.CommandUserResponse, error)
	HandleGenerateOneTimePassword(generateOtpDto *userDto.GenerateOtpDto, externalGormTransaction *gorm.DB)
	HandleVerifyOneTimePassword(verifyOtpDto *userDto.VerifyOtpDto)
	HandleGoogleAuthentication() string
	HandleGoogleCallback(tokenState string, queryCode string) *exception.ClientError
	HandleLogin(context.Context, *user.LoginUserRequest) (*user.PayloadResponse, error)
}
