package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
}

func NewRepository() *CategoryServiceImpl {
	return &CategoryServiceImpl{}
}

func (categoryRepository *CategoryServiceImpl) FindAll(gormTransaction *gorm.DB) []model.Category {
	var categoriesModel []model.Category
	err := gormTransaction.
		Preload("Jobs").
		Joins("JOIN jobs ON categories.id = jobs.category_id").
		Find(&categoriesModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return categoriesModel
}

func (categoryRepository *CategoryServiceImpl) IsCategoryExists(categoryId uint64, gormTransaction *gorm.DB) bool {
	var isCategoryExists bool
	gormTransaction.Model(&model.Category{}).
		Select("COUNT(*) > 0").
		Where("id = ?", categoryId).First(&isCategoryExists)
	return isCategoryExists
}
