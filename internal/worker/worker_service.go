package worker

import "go-takemikazuchi-api/internal/worker/dto"
import userDto "go-takemikazuchi-api/internal/user/dto"
import workerResourceDto "go-takemikazuchi-api/internal/worker_resource/dto"

type Service interface {
	Create(userJwtClaim *userDto.JwtClaimDto, createWorkerDto *dto.CreateWorkerDto, createWorkerWalletDto *workerResourceDto.CreateWorkerWalletDocumentDto)
}
