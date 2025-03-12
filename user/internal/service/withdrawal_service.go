package service

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
)

type WithdrawalService interface {
	Create(ctx context.Context, userJwtClaims *userDto.JwtClaimDto, createWithdrawalDto *dto.CreateWithdrawalDto)
	FindAll(userJwtClaims *userDto.JwtClaimDto) []model.Withdrawal
	Update(claim *userDto.JwtClaimDto, withdrawalId *string)
}
