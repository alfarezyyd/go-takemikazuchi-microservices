package worker_wallet

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Store(gormTransaction *gorm.DB, workerWalletModel *model.WorkerWallet)
}
