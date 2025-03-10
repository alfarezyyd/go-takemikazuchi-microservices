package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type JobRepositoryImpl struct{}

func NewJobRepository() *JobRepositoryImpl {
	return &JobRepositoryImpl{}
}

func (jobRepository *JobRepositoryImpl) FindById(gormTransaction *gorm.DB, id *uint64) (*model.Job, error) {
	var jobModel model.Job
	err := gormTransaction.Model(model.Job{}).Where("id = ?", *id).First(&jobModel).Error
	if err != nil {
		return nil, err
	}
	return &jobModel, nil
}

func (jobRepository *JobRepositoryImpl) Store(jobModel *model.Job, gormTransaction *gorm.DB) {
	err := gormTransaction.Create(jobModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobRepository *JobRepositoryImpl) Update(jobModel *model.Job, gormTransaction *gorm.DB) {
	err := gormTransaction.Debug().Where("id = ?", jobModel.ID).Updates(&jobModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobRepository *JobRepositoryImpl) Delete(jobId string, userId uint64, gormTransaction *gorm.DB) {
	err := gormTransaction.Joins("Users").Where("id = ? AND users.id = ?", jobId, userId).Delete(&model.Job{}).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobRepository *JobRepositoryImpl) IsExists(jobId uint64, gormTransaction *gorm.DB) bool {
	var isJobExists bool
	gormTransaction.Model(&model.Job{}).
		Select("COUNT(*) > 0").
		Where("id = ?", jobId).First(&isJobExists)
	return isJobExists
}

func (jobRepository *JobRepositoryImpl) VerifyJobOwner(gormTransaction *gorm.DB, userEmail *string, jobId *uint64) (bool, error) {
	var isJobValid bool
	err := gormTransaction.Model(&model.Job{}).
		Joins("JOIN users ON users.id = jobs.user_id").
		Select("COUNT(*) > 0").
		Where("jobs.id = ? AND users.email = ?", jobId, userEmail).
		First(&isJobValid).Error
	return isJobValid, err

}

func (jobRepository *JobRepositoryImpl) VerifyJobWorker(gormTransaction *gorm.DB, workerId *string, jobId *uint64) (bool, error) {
	var isJobValid bool
	err := gormTransaction.Model(&model.Job{}).
		Select("COUNT(*) > 0").
		Where("jobs.id = ? AND workers.id = ?", jobId, workerId).
		First(&isJobValid).Error
	return isJobValid, err
}

func (jobRepository *JobRepositoryImpl) FindVerifyById(gormTransaction *gorm.DB, userId *uint64, jobId *uint64) (*model.Job, error) {
	var jobModel model.Job
	err := gormTransaction.Model(&model.Job{}).
		Select("*").
		Where("id = ? AND user_id = ?", jobId, userId).
		First(&jobModel).Error
	return &jobModel, err
}

func (jobRepository *JobRepositoryImpl) FindWithRelationship(gormTransaction *gorm.DB, userEmail *string, jobId *uint64) *model.Job {
	var jobModel model.Job
	err := gormTransaction.
		Select("*").Where("jobs.id = ? AND users.email = ?", jobId, userEmail).
		First(&jobModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return &jobModel
}
