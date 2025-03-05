package worker_wallet

import (
	"go-takemikazuchi-microservices/internal/model"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
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
