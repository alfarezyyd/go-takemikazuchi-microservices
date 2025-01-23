package transaction

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(gormTransaction *gorm.DB, transactionModel *model.Transaction)
	Update(gormTransaction *gorm.DB, transactionModel *model.Transaction)
}
