package dto

import (
	workerWalletDto "go-takemikazuchi-api/internal/worker_wallet/dto"
)

type CreateWorkerDto struct {
	EmergencyPhoneNumber string `form:"emergency_phone_number" validate:"required,min=8,max=32"`
	Location             string `form:"location" validate:"required,min=1,max=200"`
	WalletInformation    workerWalletDto.CreateWorkerWalletDto
}
