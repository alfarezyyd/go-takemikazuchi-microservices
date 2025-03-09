package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type WorkerRepositoryImpl struct {
}

func NewWorkerRepository() *WorkerRepositoryImpl {
	return &WorkerRepositoryImpl{}
}

func (workerRepository *WorkerRepositoryImpl) Store(gormTransaction *gorm.DB, workerModel *model.Worker) {
	err := gormTransaction.Create(workerModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (workerRepository *WorkerRepositoryImpl) FindById(gormTransaction *gorm.DB, userId *uint64) (*model.Worker, error) {
	var workerModel model.Worker
	err := gormTransaction.Where("user_id = ?", userId).First(&workerModel).Error
	return &workerModel, err
}

func (workerRepository *WorkerRepositoryImpl) DynamicUpdate(gormTransaction *gorm.DB, whereClause interface{}, updatedValue interface{}, whereArgument ...interface{}) {
	err := gormTransaction.Model(&model.Worker{}).Debug().Where(whereClause, whereArgument).Updates(updatedValue).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
