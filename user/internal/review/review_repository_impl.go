package review

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

func (reviewRepository *RepositoryImpl) Create(gormTransaction *gorm.DB, reviewModel *model.Review) {
	err := gormTransaction.Create(reviewModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
