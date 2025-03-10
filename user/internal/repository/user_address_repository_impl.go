package repository

import (
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type UserAddressRepositoryImpl struct {
}

func NewUserAddressRepository() *UserAddressRepositoryImpl {
	return &UserAddressRepositoryImpl{}
}

func (userAddressRepository *UserAddressRepositoryImpl) FindById(gormTransaction *gorm.DB, id *uint64, userAddress *model.UserAddress) {
	err := gormTransaction.Where("id = ?", id).First(userAddress).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (userAddressRepository *UserAddressRepositoryImpl) Store(gormTransaction *gorm.DB, userAddress *model.UserAddress) {
	fmt.Println(userAddress)
	err := gormTransaction.Create(userAddress).Error
	fmt.Print(err)
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
