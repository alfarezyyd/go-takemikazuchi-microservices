package user

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindUserByEmail(userEmail *string, userModel *model.User, gormConnection *gorm.DB)
}
