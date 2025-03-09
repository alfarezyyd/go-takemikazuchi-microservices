package repository

import (
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
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

func (jobResourceRepository *RepositoryImpl) BulkCreate(gormTransaction *gorm.DB, jobResourceModels []*model.JobResource) {
	err := gormTransaction.Create(&jobResourceModels).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobResourceRepository *RepositoryImpl) CountBulkByName(gormTransaction *gorm.DB, jobId uint64, deletedFilesName []string) int {
	var countFile int
	err := gormTransaction.Model(&model.JobResource{}).Select("COUNT(*)").Where("job_id = ? AND image_path IN (?)", jobId, deletedFilesName).First(&countFile).Error
	fmt.Println(err)
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return countFile
}
func (jobResourceRepository *RepositoryImpl) DeleteBulkByName(gormTransaction *gorm.DB, jobId uint64, deletedFilesName []string) {
	err := gormTransaction.Model(&model.JobResource{}).Where("job_id = ? AND image_path IN (?)", jobId, deletedFilesName).Delete(&model.JobResource{}).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobResourceRepository *RepositoryImpl) DeleteBulkByJobId(gormTransaction *gorm.DB, jobId *uint64) {
	err := gormTransaction.Model(&model.JobResource{}).
		Where("job_id = ?", jobId).
		Delete(&model.JobResource{}).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
