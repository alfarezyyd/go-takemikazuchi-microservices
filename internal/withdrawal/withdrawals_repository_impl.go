package withdrawal

import (
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct{}

func NewRepositoryImpl() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (withdrawalRepository *RepositoryImpl) Create(gormTransaction *gorm.DB, withdrawalModel *model.Withdrawal) {
	err := gormTransaction.Create(withdrawalModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
