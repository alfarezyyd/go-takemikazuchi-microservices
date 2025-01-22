package user_address

import (
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewUserAddressRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (userAddressRepository *RepositoryImpl) Store(gormTransaction *gorm.DB, userAddress *model.UserAddress) {
	err := gormTransaction.Create(userAddress).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
