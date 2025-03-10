package service

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
)

type WithdrawalService interface {
	Create(userJwtClaims *userDto.JwtClaimDto, createWithdrawalDto *userDto.CreateWithdrawalDto)
	FindAll(userJwtClaims *userDto.JwtClaimDto) []model.Withdrawal
	Update(claim *userDto.JwtClaimDto, withdrawalId *string)
}
