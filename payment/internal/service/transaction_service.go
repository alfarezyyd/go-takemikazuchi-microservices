package service

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
)

type TransactionService interface {
	Create(ctx context.Context, userJwtClaims *userDto.JwtClaimDto, createTransactionDto *dto.CreateTransactionDto) string
	PostPayment(transactionNotificationDto *dto.TransactionNotificationDto)
}
