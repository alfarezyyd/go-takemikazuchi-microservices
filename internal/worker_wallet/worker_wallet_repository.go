package worker_wallet

import (
	"go-takemikazuchi-microservices/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindById(gormTransaction *gorm.DB, walletId *uint64)
	Store(gormTransaction *gorm.DB, workerWalletModel *model.WorkerWallet)
}
