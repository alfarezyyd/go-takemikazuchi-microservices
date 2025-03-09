package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type WorkerWalletRepositoryImpl struct {
}

func NewWorkerWalletRepository() *WorkerWalletRepositoryImpl {
	return &WorkerWalletRepositoryImpl{}
}

func (workerWalletRepository *WorkerWalletRepositoryImpl) FindById(gormTransaction *gorm.DB, walletId *uint64) {
	var workerWalletModel model.WorkerWallet
	err := gormTransaction.Where("id = ?", walletId).First(&workerWalletModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (workerWalletRepository *WorkerWalletRepositoryImpl) Store(gormTransaction *gorm.DB, workerWalletModel *model.WorkerWallet) {
	err := gormTransaction.Create(workerWalletModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
