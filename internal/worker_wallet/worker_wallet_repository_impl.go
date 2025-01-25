package worker_wallet

import (
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (workerWalletRepository *RepositoryImpl) FindById(gormTransaction *gorm.DB, walletId *uint64) {
	var workerWalletModel model.WorkerWallet
	err := gormTransaction.Where("id = ?", walletId).First(&workerWalletModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (workerWalletRepository *RepositoryImpl) Store(gormTransaction *gorm.DB, workerWalletModel *model.WorkerWallet) {
	err := gormTransaction.Create(workerWalletModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
