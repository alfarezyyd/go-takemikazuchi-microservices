package user_address

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindById(gormTransaction *gorm.DB, id *uint64, userAddress *model.UserAddress)
	Store(gormTransaction *gorm.DB, userAddress *model.UserAddress)
}
