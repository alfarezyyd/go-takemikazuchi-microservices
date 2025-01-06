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
