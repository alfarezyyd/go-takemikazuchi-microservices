package model

import "time"

type Withdrawal struct {
	ID              uint64        `gorm:"column:id;primaryKey;autoIncrement"`
	WorkerId        uint64        `gorm:"column:worker_id;"`
	WalletId        uint64        `gorm:"column:wallet_id;"`
	AdminId         uint64        `gorm:"column:admin_id"`
	Amount          float64       `gorm:"column:amount"`
	Status          string        `gorm:"column:status"`
	RejectionReason string        `gorm:"column:rejection_reason"`
	RequestedAt     time.Time     `gorm:"column:requested_at"`
	ProcessedAt     *time.Time    `gorm:"column:processed_at"`
	Worker          *Worker       `gorm:"foreignKey:worker_id;references:id"`
	WorkerWallet    *WorkerWallet `gorm:"foreignKey:wallet_id;references:id"`
	Admin           *User         `gorm:"foreignKey:admin_id;references:id"`
}
