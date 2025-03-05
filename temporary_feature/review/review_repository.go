package review

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(gormTransaction *gorm.DB, reviewModel *model.Review)
}
