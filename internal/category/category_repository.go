package category

import "gorm.io/gorm"

type Repository interface {
	IsCategoryExists(categoryId uint64, gormTransaction *gorm.DB) bool
}
