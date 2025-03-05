package model

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-job/internal/model"
	model2 "github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/model"
	"time"
)

type Transaction struct {
	ID            string       `gorm:"column:id;primaryKey"`
	JobID         uint64       `gorm:"column:job_id"`
	PayerID       uint64       `gorm:"column:payer_id"`
	PayeeID       uint64       `gorm:"column:payee_id"`
	Amount        float64      `gorm:"column:amount"`
	SnapToken     *string      `gorm:"column:snap_token"`
	PaymentMethod *string      `gorm:"column:payment_method"`
	Status        string       `gorm:"column:status;default:Pending"`
	CreatedAt     *time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     *time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Job           *model.Job   `gorm:"foreignKey:job_id;references:id"`
	PayerUser     *model2.User `gorm:"foreignKey:payer_id;references:id"`
	PayeeUser     *model2.User `gorm:"foreignKey:payee_id;references:id"`
}
