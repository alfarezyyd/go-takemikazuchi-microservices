package transaction

import (
	"go-takemikazuchi-microservices/internal/transaction/dto"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
)

type Service interface {
	Create(userJwtClaims *userDto.JwtClaimDto, createTransactionDto *dto.CreateTransactionDto) string
	PostPayment(transactionNotificationDto *dto.TransactionNotificationDto)
}
