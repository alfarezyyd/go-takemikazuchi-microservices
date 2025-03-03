package user

import (
	"fmt"
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

func (userRepository *RepositoryImpl) IsUserExists(gormTransaction *gorm.DB, queryClause string, argumentClause ...interface{}) (bool, error) {
	var isUserExists bool
	fmt.Println(queryClause, argumentClause)
	err := gormTransaction.Model(model.User{}).
		Select("COUNT(*) > 0").
		Where(queryClause, argumentClause...).
		First(&isUserExists).Error
	return isUserExists, err
}

func (userRepository *RepositoryImpl) FindUserByEmail(userEmail *string, userModel *model.User, gormConnection *gorm.DB) {
	err := gormConnection.Where("email = ?", userEmail).First(userModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
