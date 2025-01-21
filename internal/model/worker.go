package model

import "time"

type Worker struct {
	ID                   uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserId               uint64     `gorm:"column:user_id"`
	Rating               float32    `gorm:"column:rating;default:0"`
	Revenue              uint32     `gorm:"column:revenue;default:0"`
	CompletedJobs        uint32     `gorm:"column:completed_jobs;default:0"`
	Location             string     `gorm:"column:location"`
	Availability         bool       `gorm:"column:availability;default:true"`
	Verified             bool       `gorm:"column:verified;default:false"`
	EmergencyPhoneNumber string     `gorm:"column:emergency_phone_number" mapstructure:"emergency_phone_number"`
	CreatedAt            time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time  `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	VerifiedAt           *time.Time `gorm:"column:verified_at;"`
}
