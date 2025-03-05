package model

import (
	userModel "github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/model"
	"gorm.io/gorm"
	"time"
)

type Job struct {
	ID             uint64                 `gorm:"column:id;autoIncrement;primaryKey"`
	UserId         uint64                 `gorm:"column:user_id"`
	AddressId      uint64                 `gorm:"column:address_id"`
	CategoryId     uint64                 `gorm:"column:category_id" mapstructure:"CategoryId"`
	WorkerId       uint64                 `gorm:"column:worker_id" mapstructure:"WorkerId"`
	Title          string                 `gorm:"column:title" mapstructure:"Title"`
	Description    string                 `gorm:"column:description" mapstructure:"Description"`
	Price          float64                `gorm:"column:price" mapstructure:"Price"`
	Status         string                 `gorm:"column:status;default:'Open'"`
	CreatedAt      *time.Time             `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      *time.Time             `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User           *userModel.User        `gorm:"foreignKey:user_id;references:id"`
	Transaction    *model.Transaction     `gorm:"foreignKey:job_id;references:id"`
	UserAddress    *userModel.UserAddress `gorm:"foreignKey:address_id;references:id"`
	Category       *model.Category        `gorm:"foreignKey:category_id;references:id"`
	JobApplication []model.JobApplication `gorm:"foreignKey:job_id;references:id"`
	Worker         *model.Worker          `gorm:"foreignKey:worker_id;references:id"`
}

func (jobModel *Job) BeforeUpdate(tx *gorm.DB) (err error) {
	timeNow := time.Now()
	jobModel.UpdatedAt = &timeNow
	return nil
}
