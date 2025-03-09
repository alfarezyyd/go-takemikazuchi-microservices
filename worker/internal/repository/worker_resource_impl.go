package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type WorkerResourceRepositoryImpl struct {
}

func NewRepository() *WorkerResourceRepositoryImpl {
	return &WorkerResourceRepositoryImpl{}
}

func (workerResourceRepository *WorkerResourceRepositoryImpl) BulkStore(gormTransaction *gorm.DB, workerResourcesModel []*model.WorkerResource) {
	err := gormTransaction.Create(&workerResourcesModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))

}
