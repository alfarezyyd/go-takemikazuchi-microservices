package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type JobApplicationRepository interface {
	BulkRejectUpdate(gormTransaction *gorm.DB, jobId *uint64)
	Update(gormTransaction *gorm.DB, jobApplicationModel *model.JobApplication)
	FindAllApplication(gormTransaction *gorm.DB, jobId *uint64) []model.JobApplication
	FindById(gormTransaction *gorm.DB, userId *uint64, jobId *uint64) *model.JobApplication
}
