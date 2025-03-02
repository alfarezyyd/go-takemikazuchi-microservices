package withdrawal

import (
	"go-takemikazuchi-microservices/internal/model"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct{}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}
func (withdrawalRepository *RepositoryImpl) FindAll(gormTransaction *gorm.DB) []model.Withdrawal {
	var withdrawalsModel []model.Withdrawal
	err := gormTransaction.Model(&model.Withdrawal{}).
		Preload("Worker").
		Preload("WorkerWallet").
		Joins("JOIN worker_wallets ON worker_wallets.id = withdrawals.wallet_id").
		Joins("JOIN workers ON workers.id = withdrawals.worker_id").
		Select(`withdrawals.*, worker_wallets.wallet_type, worker_wallets.account_name, worker_wallets.account_number, worker_wallets.bank_name`).
		Find(&withdrawalsModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return withdrawalsModel
}
func (withdrawalRepository *RepositoryImpl) Create(gormTransaction *gorm.DB, withdrawalModel *model.Withdrawal) {
	err := gormTransaction.Create(withdrawalModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (withdrawalRepository *RepositoryImpl) FindById(gormTransaction *gorm.DB, withdrawalId *uint64) (*model.Withdrawal, error) {
	var withdrawalsModel model.Withdrawal
	err := gormTransaction.Model(&model.Withdrawal{}).
		Preload("Worker").
		Preload("WorkerWallet").
		Joins("JOIN worker_wallets ON worker_wallets.id = withdrawals.wallet_id").
		Joins("JOIN workers ON workers.id = withdrawals.worker_id").
		Select(`withdrawals.*, worker_wallets.wallet_type, worker_wallets.account_name, worker_wallets.account_number, worker_wallets.bank_name`).
		Where("withdrawals.id = ?", withdrawalId).Find(&withdrawalsModel).
		First(&withdrawalsModel).Error
	return &withdrawalsModel, err
}
func (withdrawalRepository *RepositoryImpl) Update(gormTransaction *gorm.DB, withdrawalModel *model.Withdrawal) {
	err := gormTransaction.Where("id = ?", withdrawalModel.ID).Updates(withdrawalModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
