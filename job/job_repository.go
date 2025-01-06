package job

import (
	"go-takemikazuchi-api/model"
	"gorm.io/gorm"
)

type Repository interface {
	Store(jobModel model.Job, gormTransaction *gorm.DB)
}
