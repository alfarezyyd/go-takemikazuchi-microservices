package withdrawal

import (
	"go-takemikazuchi-api/internal/model"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/internal/withdrawal/dto"
)

type Service interface {
	Create(userJwtClaims *userDto.JwtClaimDto, createWithdrawalDto *dto.CreateWithdrawalDto)
	FindAll(userJwtClaims *userDto.JwtClaimDto) []model.Withdrawal
	Update(claim *userDto.JwtClaimDto, withdrawalId *string)
}
