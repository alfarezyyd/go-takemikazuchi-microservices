package review

import (
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepositoryImpl() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (reviewRepository *RepositoryImpl) Create(gormTransaction *gorm.DB, reviewModel *model.Review) {
	err := gormTransaction.Create(reviewModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
