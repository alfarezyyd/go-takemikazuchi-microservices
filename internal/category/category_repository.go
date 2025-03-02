package category

import (
	"go-takemikazuchi-microservices/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(gormTransaction *gorm.DB) []model.Category
	IsCategoryExists(categoryId uint64, gormTransaction *gorm.DB) bool
}
