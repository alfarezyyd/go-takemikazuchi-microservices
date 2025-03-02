package withdrawal

import (
	"go-takemikazuchi-microservices/internal/model"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
	"go-takemikazuchi-microservices/internal/withdrawal/dto"
)

type Service interface {
	Create(userJwtClaims *userDto.JwtClaimDto, createWithdrawalDto *dto.CreateWithdrawalDto)
	FindAll(userJwtClaims *userDto.JwtClaimDto) []model.Withdrawal
	Update(claim *userDto.JwtClaimDto, withdrawalId *string)
}
