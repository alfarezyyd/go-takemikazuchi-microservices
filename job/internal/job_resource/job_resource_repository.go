package job_resource

import (
	"go-takemikazuchi-microservices/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(gormTransaction *gorm.DB, jobResourceModel *model.JobResource)
	BulkCreate(gormTransaction *gorm.DB, jobResourceModels []*model.JobResource)
	CountBulkByName(gormTransaction *gorm.DB, jobId uint64, deletedFilesName []string) int
	DeleteBulkByName(gormTransaction *gorm.DB, id uint64, deletedFilesName []string)
	DeleteBulkByJobId(gormTransaction *gorm.DB, jobId *uint64)
}
