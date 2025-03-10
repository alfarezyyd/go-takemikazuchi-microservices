package repository

import (
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (userRepository *UserRepositoryImpl) IsUserExists(gormTransaction *gorm.DB, queryClause string, argumentClause ...interface{}) (bool, error) {
	var isUserExists bool
	err := gormTransaction.Model(model.User{}).
		Select("COUNT(*) > 0").
		Where(queryClause, argumentClause...).
		First(&isUserExists).Error
	return isUserExists, err
}

func (userRepository *UserRepositoryImpl) FindUserByEmail(userEmail *string, userModel *model.User, gormConnection *gorm.DB) {
	err := gormConnection.Where("email = ?", userEmail).First(userModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
