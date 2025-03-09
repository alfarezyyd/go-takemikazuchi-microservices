package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type WorkerRepository interface {
	Store(gormTransaction *gorm.DB, workerModel *model.Worker)
	FindById(gormTransaction *gorm.DB, userId *uint64) (*model.Worker, error)
	DynamicUpdate(gormTransaction *gorm.DB, whereClause interface{}, updatedValue interface{}, whereArgument ...interface{})
}
