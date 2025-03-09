package repository

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

func (transactionRepository *RepositoryImpl) FindById(gormTransaction *gorm.DB, id string) *model.Transaction {
	var transactionModel model.Transaction
	err := gormTransaction.Where("id = ?", id).First(&transactionModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return &transactionModel
}

func (transactionRepository *RepositoryImpl) FindWithRelationship(gormTransaction *gorm.DB, id string) *model.Transaction {
	var transactionModel model.Transaction
	err := gormTransaction.
		Model(&model.Transaction{}).
		Preload("PayerUser").
		Preload("Job").
		Joins("JOIN users ON users.id = transactions.payer_id").
		Joins("JOIN jobs ON jobs.id = transactions.job_id").
		Select(`
			transactions.*,
			users.id AS user_id, users.name AS user_name, users.email AS user_email,
			jobs.id AS job_id, jobs.status AS job_status
		`).Where("transactions.id = ?", id).
		First(&transactionModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return &transactionModel
}

func (transactionRepository *RepositoryImpl) Create(gormTransaction *gorm.DB, transactionModel *model.Transaction) {
	err := gormTransaction.Create(transactionModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (transactionRepository *RepositoryImpl) Update(gormTransaction *gorm.DB, transactionModel *model.Transaction) {
	err := gormTransaction.Where("id = ?", transactionModel.ID).Updates(transactionModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
