package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type WorkerWalletRepository interface {
	FindById(gormTransaction *gorm.DB, walletId *uint64) *model.WorkerWallet
	Store(gormTransaction *gorm.DB, workerWalletModel *model.WorkerWallet)
}
