package dto

type CreateTransactionDto struct {
	JobId   uint64 `json:"job_id" validate:"required,gte=1"`
	PayerId uint64 `json:"payer_id" validate:"required,gte=1"`
	PayeeId uint64 `json:"payee_id" validate:"required,gte=1"`
	Amount  uint64 `json:"amount" validate:"required,gte=1"`
}
