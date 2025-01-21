package worker

import "go-takemikazuchi-api/internal/worker/dto"
import workerResourceDto "go-takemikazuchi-api/internal/worker_resource/dto"

type Service interface {
	Create(createWorkerDto *dto.CreateWorkerDto, createWorkerWalletDto *workerResourceDto.CreateWorkerWalletDocumentDto)
}
