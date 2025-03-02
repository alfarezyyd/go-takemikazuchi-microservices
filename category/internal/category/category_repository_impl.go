package category

import (
	"go-takemikazuchi-microservices/internal/model"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (categoryRepository *RepositoryImpl) FindAll(gormTransaction *gorm.DB) []model.Category {
	var categoriesModel []model.Category
	err := gormTransaction.
		Preload("Jobs").
		Joins("JOIN jobs ON categories.id = jobs.category_id").
		Find(&categoriesModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return categoriesModel
}

func (categoryRepository *RepositoryImpl) IsCategoryExists(categoryId uint64, gormTransaction *gorm.DB) bool {
	var isCategoryExists bool
	gormTransaction.Model(&model.Category{}).
		Select("COUNT(*) > 0").
		Where("id = ?", categoryId).First(&isCategoryExists)
	return isCategoryExists
}
