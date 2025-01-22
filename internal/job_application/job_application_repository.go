package job_application

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAllApplication(gormTransaction *gorm.DB, jobId *uint64) []model.JobApplication
}
