package service

import (
	"context"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/pkg/dto"
)

type WorkerService interface {
	Create(ctx context.Context, userJwtClaim *userDto.JwtClaimDto, createWorkerDto *dto.CreateWorkerDto)
}
