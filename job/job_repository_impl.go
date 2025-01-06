package job

import (
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/helper"
	"go-takemikazuchi-api/model"
	"gorm.io/gorm"
)

type RepositoryImpl struct{}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (jobRepository *RepositoryImpl) Store(jobModel model.Job, gormTransaction *gorm.DB) {
	err := gormTransaction.Create(&jobModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobRepository *RepositoryImpl) Update(jobModel model.Job, gormTransaction *gorm.DB) {
	err := gormTransaction.Updates(&jobModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobRepository *RepositoryImpl) Delete(jobId string, userId string, gormTransaction *gorm.DB) {
	err := gormTransaction.Joins("Users").Where("id = ? AND users.id = ?", jobId, userId).Delete(&model.Job{}).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
