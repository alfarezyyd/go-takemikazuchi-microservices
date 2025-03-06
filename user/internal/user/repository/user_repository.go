package repository

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	IsUserExists(gormTransaction *gorm.DB, queryClause string, argumentClause ...interface{}) (bool, error)
	FindUserByEmail(userEmail *string, userModel *model.User, gormConnection *gorm.DB)
}
