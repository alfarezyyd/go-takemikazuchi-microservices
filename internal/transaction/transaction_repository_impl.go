package transaction

import (
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
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
		Preload("User").
		Where("id = ?", id).
		Joins("users ON users.id = transactions.user_id").
		Select(`
			transactions.*,
			users.id AS user_id, users.name AS user_name, users.email AS user_email,
		`).First(&transactionModel).Error
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
