package user

import (
	"go-takemikazuchi-api/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindUserByEmail(userEmail *string, userModel *model.User, gormConnection *gorm.DB)
}
