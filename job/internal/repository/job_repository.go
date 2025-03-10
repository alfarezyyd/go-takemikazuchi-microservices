package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type JobRepository interface {
	FindById(gormTransaction *gorm.DB, id *uint64) (*model.Job, error)
	FindVerifyById(gormTransaction *gorm.DB, userId *uint64, jobId *uint64) (*model.Job, error)
	Store(jobModel *model.Job, gormTransaction *gorm.DB)
	Update(jobModel *model.Job, gormTransaction *gorm.DB)
	Delete(jobId string, userId uint64, gormTransaction *gorm.DB)
	IsExists(jobId uint64, gormTransaction *gorm.DB) bool
	VerifyJobOwner(gormTransaction *gorm.DB, userEmail *string, jobId *uint64) (bool, error)
	VerifyJobWorker(gormTransaction *gorm.DB, userEmail *string, jobId *uint64) (bool, error)
}
