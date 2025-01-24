package review

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(gormTransaction *gorm.DB, reviewModel *model.Review)
}
