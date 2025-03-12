package dto

import "time"

type CreateWorkerDto struct {
	EmergencyPhoneNumber string `form:"emergency_phone_number" validate:"required,min=8,max=32"`
	Location             string `form:"location" validate:"required,min=1,max=200"`
	WalletInformation    CreateWorkerWalletDto
}

type WorkerResponseDto struct {
	ID                   uint64
	UserId               uint64
	Rating               float32
	Revenue              uint32
	CompletedJobs        uint32
	Location             string
	Availability         bool
	Verified             bool
	EmergencyPhoneNumber string
	CreatedAt            time.Time  `mapstructure:"-"`
	UpdatedAt            time.Time  `mapstructure:"-"`
	VerifiedAt           *time.Time `mapstructure:"-"`
}
