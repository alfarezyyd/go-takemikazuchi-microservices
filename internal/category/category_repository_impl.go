package category

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (repositoryImpl *RepositoryImpl) IsCategoryExists(categoryId uint64, gormTransaction *gorm.DB) bool {
	var isCategoryExists bool
	gormTransaction.Model(&model.Category{}).
		Select("COUNT(*) > 0").
		Where("id = ?", categoryId).First(&isCategoryExists)
	return isCategoryExists
}
