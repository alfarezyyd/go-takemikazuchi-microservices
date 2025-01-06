package user

import (
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/helper"
	"go-takemikazuchi-api/model"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (userRepository *RepositoryImpl) FindUserByEmail(userEmail *string, userModel *model.User, gormConnection *gorm.DB) {
	err := gormConnection.Where("email = ?", userEmail).First(userModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
