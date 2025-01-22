package user_address

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Store(gormTransaction *gorm.DB, userAddress *model.UserAddress)
}
