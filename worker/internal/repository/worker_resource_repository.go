package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type WorkerResourceRepository interface {
	BulkStore(gormTransaction *gorm.DB, workerResourcesModel []*model.WorkerResource)
}
