package job_resource

import (
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (jobResourceRepository *RepositoryImpl) Create(gormTransaction *gorm.DB, jobResourceModel *model.JobResource) {
	err := gormTransaction.Create(&jobResourceModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
