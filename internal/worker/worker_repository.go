package worker

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Store(gormTransaction *gorm.DB, workerModel *model.Worker)
	FindById(gormTransaction *gorm.DB, userId *uint64) (*model.Worker, error)
	DynamicUpdate(gormTransaction *gorm.DB, whereClause interface{}, updatedValue interface{}, whereArgument ...interface{})
}
