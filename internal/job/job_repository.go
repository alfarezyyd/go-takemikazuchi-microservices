package job

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Store(jobModel *model.Job, gormTransaction *gorm.DB)
	Update(jobModel model.Job, gormTransaction *gorm.DB)
	Delete(jobId string, userId uint64, gormTransaction *gorm.DB)
	IsExists(jobId uint64, gormTransaction *gorm.DB) bool
}
