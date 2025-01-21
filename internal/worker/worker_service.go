package worker

import "go-takemikazuchi-api/internal/worker/dto"

type Service interface {
	Create(createWorkerDto *dto.CreateWorkerDto, createWorkerWalletDto *dto.CreateWorkerWalletDocumentDto)
}
