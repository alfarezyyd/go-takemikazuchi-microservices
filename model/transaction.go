package model

import "time"

type Transaction struct {
	ID            uint64     `gorm:"column:id;primary_key;autoIncrement"`
	JobID         uint64     `gorm:"column:job_id"`
	PayerID       uint64     `gorm:"column:payer_id"`
	PayeeID       uint64     `gorm:"column:payee_id"`
	Amount        uint64     `gorm:"column:amount"`
	SnapToken     string     `gorm:"column:snap_token"`
	PaymentMethod string     `gorm:"column:payment_method"`
	Status        string     `gorm:"column:status;default:Pending"`
	CreatedAt     *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Job           *Job       `gorm:"foreignKey:job_id;references:job_id"`
	PayerUser     *User      `gorm:"foreignKey:payer_id;references:id"`
	PayeeUser     *User      `gorm:"foreignKey:payee_id;references:id"`
}
