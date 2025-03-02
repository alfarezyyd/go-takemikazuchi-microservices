package worker

import "go-takemikazuchi-microservices/internal/worker/dto"
import userDto "go-takemikazuchi-microservices/internal/user/dto"
import workerResourceDto "go-takemikazuchi-microservices/internal/worker_resource/dto"

type Service interface {
	Create(userJwtClaim *userDto.JwtClaimDto, createWorkerDto *dto.CreateWorkerDto, createWorkerWalletDto *workerResourceDto.CreateWorkerWalletDocumentDto)
}
