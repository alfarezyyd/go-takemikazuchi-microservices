package dto

type CreateTransactionDto struct {
	JobId       uint64 `json:"job_id" validate:"required,gte=1"`
	ApplicantId uint64 `json:"applicant_id" validate:"required,gte=1"`
}
