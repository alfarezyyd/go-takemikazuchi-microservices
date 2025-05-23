package service

import (
	"context"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/pkg/dto"
)

type WorkerService interface {
	FindById(ctx context.Context, userId *uint64) *dto.WorkerResponseDto
	Create(ctx context.Context, userJwtClaim *userDto.JwtClaimDto, createWorkerDto *dto.CreateWorkerDto)
}
