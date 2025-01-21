package worker_resource

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (workerResourceRepository *RepositoryImpl) BulkStore(gormTransaction *gorm.DB, workerResourcesModel []*model.WorkerResource) {
}
