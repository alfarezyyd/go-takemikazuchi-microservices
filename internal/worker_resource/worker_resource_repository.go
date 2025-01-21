package worker_resource

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	BulkStore(gormTransaction *gorm.DB, workerResourcesModel []*model.WorkerResource)
}
