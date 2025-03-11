package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct{}

func NewTransactionRepository() *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{}

}

func (transactionRepository *TransactionRepositoryImpl) FindById(gormTransaction *gorm.DB, id string) *model.Transaction {
	var transactionModel model.Transaction
	err := gormTransaction.Where("id = ?", id).First(&transactionModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return &transactionModel
}

func (transactionRepository *TransactionRepositoryImpl) FindWithRelationship(gormTransaction *gorm.DB, id string) *model.Transaction {
	var transactionModel model.Transaction
	err := gormTransaction.
		Model(&model.Transaction{}).
		Select(`*`).Where("transactions.id = ?", id).
		First(&transactionModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return &transactionModel
}

func (transactionRepository *TransactionRepositoryImpl) Create(gormTransaction *gorm.DB, transactionModel *model.Transaction) {
	err := gormTransaction.Create(transactionModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (transactionRepository *TransactionRepositoryImpl) Update(gormTransaction *gorm.DB, transactionModel *model.Transaction) {
	err := gormTransaction.Where("id = ?", transactionModel.ID).Updates(transactionModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
