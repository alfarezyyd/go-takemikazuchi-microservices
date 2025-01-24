package model

import "time"

type Review struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	ReviewerId uint64    `gorm:"column:reviewer_id"`
	ReviewedId uint64    `gorm:"column:reviewed_id" mapstructure:Revie`
	JobId      uint64    `gorm:"column:job_id"`
	Role       string    `gorm:"column:role"`
	Rating     byte      `gorm:"column:rating"`
	ReviewText string    `gorm:"column:review_text"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Reviewer   *User     `gorm:"foreignKey:reviewer_id;references:id"`
	Reviewed   *User     `gorm:"foreignKey:reviewed_id;references:id"`
	Job        *Job      `gorm:"foreignKey:job_id;references:id"`
}
