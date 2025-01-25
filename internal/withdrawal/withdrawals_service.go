package withdrawal

import (
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/internal/withdrawal/dto"
)

type Service interface {
	Create(userJwtClaims *userDto.JwtClaimDto, createWithdrawalDto *dto.CreateWithdrawalDto)
}
