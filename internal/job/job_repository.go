package job

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindById(gormTransaction *gorm.DB, id *uint64) (*model.Job, error)
	FindVerifyById(gormTransaction *gorm.DB, id *uint64, userEmail *string) (*model.Job, error)
	Store(jobModel *model.Job, gormTransaction *gorm.DB)
	Update(jobModel *model.Job, gormTransaction *gorm.DB)
	Delete(jobId string, userId uint64, gormTransaction *gorm.DB)
	IsExists(jobId uint64, gormTransaction *gorm.DB) bool
	VerifyJobOwner(gormTransaction *gorm.DB, userEmail *string, jobId *uint64)
}
