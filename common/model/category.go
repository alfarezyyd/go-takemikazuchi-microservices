package model

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-job/internal/model"
	"time"
)

type Category struct {
	ID          uint64      `gorm:"primary_key;autoIncrement"`
	Name        string      `gorm:"column:name" mapstructure:"name"`
	Description string      `gorm:"column:description" mapstructure:"description"`
	CreatedAt   time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Jobs        []model.Job `gorm:"foreignKey:category_id;references:id"`
}
