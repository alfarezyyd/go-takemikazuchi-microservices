package transaction

import (
	"go-takemikazuchi-api/internal/transaction/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
)

type Service interface {
	Create(userJwtClaims *userDto.JwtClaimDto, createTransactionDto *dto.CreateTransactionDto) string
}
