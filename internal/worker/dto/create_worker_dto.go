package dto

import (
	workerWalletDto "go-takemikazuchi-api/internal/worker_wallet/dto"
	"mime/multipart"
)

type CreateWorkerDto struct {
	EmergencyPhoneNumber string `json:"emergency_phone_number" validate:"required;min=8,max=32"`
	Location             string `json:"location" validate:"required;min=1,max=200"`
	WalletInformation    workerWalletDto.CreateWorkerWalletDto
}
type CreateWorkerWalletDocumentDto struct {
	IdentityCard      *multipart.FileHeader
	PoliceCertificate *multipart.FileHeader
	DriverLicense     *multipart.FileHeader
}
