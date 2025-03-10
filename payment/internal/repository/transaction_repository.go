package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindById(gormTransaction *gorm.DB, id string) *model.Transaction
	Create(gormTransaction *gorm.DB, transactionModel *model.Transaction)
	Update(gormTransaction *gorm.DB, transactionModel *model.Transaction)
	FindWithRelationship(gormTransaction *gorm.DB, id string) *model.Transaction
}
