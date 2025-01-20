package model

import "time"

type Worker struct {
	ID            uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserId        uint64     `gorm:"column:user_id"`
	Rating        float32    `gorm:"column:rating"`
	CompletedJobs uint32     `gorm:"column:completed_jobs"`
	Location      string     `gorm:"column:location"`
	Availability  bool       `gorm:"column:availability;default:true"`
	Verified      bool       `gorm:"column:verified;default:false"`
	CreatedAt     time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	VerifiedAt    *time.Time `gorm:"column:verified_at;"`
}
