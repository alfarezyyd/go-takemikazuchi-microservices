package worker_resource

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

func (workerResourceRepository *RepositoryImpl) BulkStore(gormTransaction *gorm.DB, workerResourcesModel []*model.WorkerResource) {
	err := gormTransaction.Create(&workerResourcesModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))

}
