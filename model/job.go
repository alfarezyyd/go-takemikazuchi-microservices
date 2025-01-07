package model

import "time"

type Job struct {
	ID          uint64       `gorm:"column:id;autoIncrement;primaryKey"`
	UserId      uint64       `gorm:"column:user_id"`
	CategoryId  uint64       `gorm:"column:category_id" mapstructure:"CategoryId"`
	Title       string       `gorm:"column:title" mapstructure:"Title"`
	Description string       `gorm:"column:description" mapstructure:"Description"`
	Latitude    float64      `gorm:"column:latitude"`
	Longitude   float64      `gorm:"column:longitude"`
	Address     string       `gorm:"column:address"`
	PlaceId     string       `gorm:"column:place_id"`
	Price       float64      `gorm:"column:price" mapstructure:"Price"`
	Status      string       `gorm:"column:status;default:'Open'"`
	CreatedAt   *time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   *time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Transaction *Transaction `gorm:"foreignKey:job_id;references:id"`
}
