package user

import (
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (userRepository *RepositoryImpl) IsUserExists(gormTransaction *gorm.DB, userId *uint64, userEmail *string) (bool, error) {
	var isUserExists bool
	err := gormTransaction.Model(model.User{}).
		Select("COUNT(*) > 0").
		Where("id = ? OR email = ?", userId, userEmail).
		First(&isUserExists).Error
	return isUserExists, err
}

func (userRepository *RepositoryImpl) FindUserByEmail(userEmail *string, userModel *model.User, gormConnection *gorm.DB) {
	err := gormConnection.Where("email = ?", userEmail).First(userModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
