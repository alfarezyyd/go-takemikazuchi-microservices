package model

import "time"

type JobApplication struct {
	ID          uint64     `gorm:"column:id;primaryKey;autoIncrement"`
	JobId       uint64     `gorm:"column:job_id"`
	ApplicantId uint64     `gorm:"column:applicant_id"`
	Status      string     `gorm:"column:status;default:Pending"`
	CreatedAt   *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Job         *Job       `gorm:"foreignKey:job_id;references:id"`
	User        *User      `gorm:"foreignKey:applicant_id;references:id"`
}
