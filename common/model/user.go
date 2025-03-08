package model

import (
	"time"
)

type User struct {
	ID              uint64     `gorm:"column:id;primary_key;autoIncrement"`
	Name            string     `gorm:"column:name" mapstructure:"name"`
	Email           string     `gorm:"column:email" mapstructure:"email"`
	Password        string     `gorm:"column:password" mapstructure:"password"`
	Role            string     `gorm:"column:role;default:Employer"`
	PhoneNumber     *string    `gorm:"column:phone_number" mapstructure:"phone_number"`
	ProfilePicture  *string    `gorm:"column:profile_picture"`
	IsActive        bool       `gorm:"column:is_active;default:true"`
	CreatedAt       time.Time  `gorm:"column:created_at;autoCreateTime" mapstructure:"-"`
	UpdatedAt       time.Time  `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" mapstructure:"-"`
	EmailVerifiedAt *time.Time `gorm:"column:email_verified_at"`
}
