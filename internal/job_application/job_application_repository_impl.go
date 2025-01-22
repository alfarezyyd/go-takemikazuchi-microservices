package job_application

import (
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct{}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (jobApplicationRepository RepositoryImpl) FindAllApplication(gormTransaction *gorm.DB, jobId *uint64) []model.JobApplication {
	var jobApplications []model.JobApplication
	err := gormTransaction.
		Preload("Job").
		Preload("User").
		Joins("JOIN jobs ON jobs.id = job_applications.job_id").
		Joins("JOIN users ON users.id = job_applications.applicant_id").
		Select("job_applications.*, jobs.id as job_id, users.id as applicant_id").
		Where("jobs.id = ?", jobId).
		Find(&jobApplications).Error

	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return jobApplications
}
