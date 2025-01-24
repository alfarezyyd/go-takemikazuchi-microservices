package worker

import (
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (workerRepository *RepositoryImpl) Store(gormTransaction *gorm.DB, workerModel *model.Worker) {
	err := gormTransaction.Create(workerModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (workerRepository *RepositoryImpl) FindById(gormTransaction *gorm.DB, userId *uint64) (*model.Worker, error) {
	var workerModel model.Worker
	err := gormTransaction.Where("user_id = ?", userId).First(&workerModel).Error
	return &workerModel, err
}
