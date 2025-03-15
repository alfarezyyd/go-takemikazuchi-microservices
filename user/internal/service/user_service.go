package service

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"gorm.io/gorm"
)

type UserService interface {
	HandleRegister(createUserDto *dto.CreateUserDto) error
	HandleGenerateOneTimePassword(generateOtpDto *dto.GenerateOtpDto, externalGormTransaction *gorm.DB)
	HandleVerifyOneTimePassword(verifyOtpDto *dto.VerifyOtpDto)
	HandleGoogleAuthentication() string
	HandleGoogleCallback(tokenState string, queryCode string) *exception.ClientError
	HandleLogin(ctx context.Context, loginUserDto *dto.LoginUserDto) string
	FindByIdentifier(ctx context.Context, userIdentifierDto *dto.UserIdentifierDto) *dto.UserResponseDto
}
