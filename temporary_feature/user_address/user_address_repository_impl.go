package user_address

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/model"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewUserAddressRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (userAddressRepository *RepositoryImpl) FindById(gormTransaction *gorm.DB, id *uint64, userAddress *model.UserAddress) {
	err := gormTransaction.Where("id = ?", id).First(userAddress).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (userAddressRepository *RepositoryImpl) Store(gormTransaction *gorm.DB, userAddress *model.UserAddress) {
	err := gormTransaction.Create(userAddress).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
