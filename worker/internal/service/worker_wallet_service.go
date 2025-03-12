package service

import (
	"context"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/pkg/dto"
)

type WorkerWalletService interface {
	FindById(ctx context.Context, walletId *uint64) *dto.ResponseWorkerWalletDto
	Create(ctx context.Context, userJwtClaim *userDto.JwtClaimDto, createWorkerDto *dto.CreateWorkerDto)
}
