package user

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	IsUserExists(gormTransaction *gorm.DB, userId *uint64, userEmail *string) (bool, error)
	FindUserByEmail(userEmail *string, userModel *model.User, gormConnection *gorm.DB)
}
