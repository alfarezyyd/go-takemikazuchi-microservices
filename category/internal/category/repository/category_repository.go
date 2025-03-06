package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(gormTransaction *gorm.DB) []model.Category
	IsCategoryExists(categoryId uint64, gormTransaction *gorm.DB) bool
}
