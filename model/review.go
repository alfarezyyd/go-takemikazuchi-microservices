package model

import "time"

type Review struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	ReviewerId uint64    `gorm:"column:reviewer_id"`
	RevieweeId uint64    `gorm:"column:reviewee_id"`
	Rating     uint64    `gorm:"column:rating"`
	Comment    string    `gorm:"column:comment"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
