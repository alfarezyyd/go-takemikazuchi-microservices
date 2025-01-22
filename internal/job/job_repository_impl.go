package job

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

func (jobRepository *RepositoryImpl) FindById(gormTransaction *gorm.DB, id *uint64) (*model.Job, error) {
	var jobModel model.Job
	err := gormTransaction.Model(model.Job{}).Where("id = ?", *id).First(&jobModel).Error
	if err != nil {
		return nil, err
	}
	return &jobModel, nil
}

func (jobRepository *RepositoryImpl) Store(jobModel *model.Job, gormTransaction *gorm.DB) {
	err := gormTransaction.Create(jobModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobRepository *RepositoryImpl) Update(jobModel *model.Job, gormTransaction *gorm.DB) {
	err := gormTransaction.Updates(&jobModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobRepository *RepositoryImpl) Delete(jobId string, userId uint64, gormTransaction *gorm.DB) {
	err := gormTransaction.Joins("Users").Where("id = ? AND users.id = ?", jobId, userId).Delete(&model.Job{}).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobRepository *RepositoryImpl) IsExists(jobId uint64, gormTransaction *gorm.DB) bool {
	var isJobExists bool
	gormTransaction.Model(&model.Job{}).
		Select("COUNT(*) > 0").
		Where("id = ?", jobId).First(&isJobExists)
	return isJobExists
}

func (jobRepository *RepositoryImpl) VerifyJobOwner(gormTransaction *gorm.DB, userEmail *string, jobId *uint64) {
	var isJobValid bool
	err := gormTransaction.Model(&model.Job{}).
		Joins("JOIN users ON users.id = jobs.user_id").
		Select("COUNT(*) > 0").
		Where("jobs.id = ? AND users.email = ?", jobId, userEmail).
		First(&isJobValid).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))

}

func (jobRepository *RepositoryImpl) FindVerifyById(gormTransaction *gorm.DB, userEmail *string, jobId *uint64) *model.Job {
	var jobModel model.Job
	err := gormTransaction.Model(&model.Job{}).
		Joins("JOIN users ON users.id = jobs.user_id").
		Select("*").
		Where("jobs.id = ? AND users.email = ?", jobId, userEmail).
		First(&jobModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return &jobModel
}
