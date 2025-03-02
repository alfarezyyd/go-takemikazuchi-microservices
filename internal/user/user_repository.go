package user

import (
	"go-takemikazuchi-microservices/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	IsUserExists(gormTransaction *gorm.DB, queryClause string, argumentClause ...interface{}) (bool, error)
	FindUserByEmail(userEmail *string, userModel *model.User, gormConnection *gorm.DB)
}
