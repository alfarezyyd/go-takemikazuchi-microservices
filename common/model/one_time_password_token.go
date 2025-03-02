package model

import "time"

type OneTimePasswordToken struct {
	ID          uint64    `gorm:"column:id;primary_key;autoIncrement"`
	UserId      uint64    `gorm:"column:user_id"`
	User        User      `gorm:"foreignKey:user_id;references:id"`
	HashedToken string    `gorm:"column:hashed_token"`
	ExpiresAt   time.Time `gorm:"column:expires_at"`
}
