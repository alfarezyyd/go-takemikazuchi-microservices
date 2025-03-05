package review

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/model"
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
