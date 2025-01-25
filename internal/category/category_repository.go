package category

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(gormTransaction *gorm.DB) []model.Category
	IsCategoryExists(categoryId uint64, gormTransaction *gorm.DB) bool
}
