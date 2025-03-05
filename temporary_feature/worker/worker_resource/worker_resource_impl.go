package worker_resource

import (
	"go-takemikazuchi-microservices/internal/model"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (workerResourceRepository *RepositoryImpl) BulkStore(gormTransaction *gorm.DB, workerResourcesModel []*model.WorkerResource) {
	err := gormTransaction.Create(&workerResourcesModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))

}
