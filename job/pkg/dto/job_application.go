package dto

type ApplyJobApplicationDto struct {
	JobId uint64 `json:"job_id" validate:"required,gte=1"`
}

type JobApplicationResponseDto struct {
	Id          uint64 `json:"id"`
	FullName    string `json:"full_name"`
	AppliedAt   string `json:"applied_at"`
	JobId       uint64 `json:"job_id"`
	ApplicantId uint64 `json:"applicant_id"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type SelectApplicationDto struct {
	UserId uint64 `json:"user_id"`
	JobId  uint64 `json:"job_id"`
}
