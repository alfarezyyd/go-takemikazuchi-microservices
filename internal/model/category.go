package model

import "time"

type Category struct {
	ID          uint64    `gorm:"primary_key;autoIncrement"`
	Name        string    `gorm:"column:name" mapstructure:"name"`
	Description string    `gorm:"column:description" mapstructure:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
