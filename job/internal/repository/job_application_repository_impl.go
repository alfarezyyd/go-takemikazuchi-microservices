package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type JobApplicationRepositoryImpl struct{}

func NewJobApplicationRepository() *JobApplicationRepositoryImpl {
	return &JobApplicationRepositoryImpl{}
}

func (jobApplicationRepository *JobApplicationRepositoryImpl) FindAllApplication(gormTransaction *gorm.DB, jobId *uint64) []model.JobApplication {
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

func (jobApplicationRepository *JobApplicationRepositoryImpl) BulkRejectUpdate(gormTransaction *gorm.DB, jobId *uint64) {
	err := gormTransaction.Where("job_id = ?", jobId).Updates(model.JobApplication{
		Status: "Rejected",
	}).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobApplicationRepository *JobApplicationRepositoryImpl) FindById(gormTransaction *gorm.DB, userId *uint64, jobId *uint64) *model.JobApplication {
	var jobApplication model.JobApplication
	err := gormTransaction.
		Where("job_id = ? AND applicant_id = ?", jobId, userId).
		First(&jobApplication).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return &jobApplication
}

func (jobApplicationRepository *JobApplicationRepositoryImpl) Update(gormTransaction *gorm.DB, jobApplicationModel *model.JobApplication) {
	err := gormTransaction.
		Where("id = ?", jobApplicationModel.ID).
		Updates(jobApplicationModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
