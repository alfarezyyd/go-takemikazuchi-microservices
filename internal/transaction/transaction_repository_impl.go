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

func (transactionRepository *RepositoryImpl) Create(gormTransaction *gorm.DB, transactionModel *model.Transaction) {
	err := gormTransaction.Create(transactionModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (transactionRepository *RepositoryImpl) Update(gormTransaction *gorm.DB, transactionModel *model.Transaction) {
	err := gormTransaction.Where("id = ?", transactionModel.ID).Updates(transactionModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
