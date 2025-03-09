package dto

type ApplyJobApplicationDto struct {
	JobId uint64 `json:"job_id" validate:"required,gte=1"`
}

type JobApplicationResponseDto struct {
	Id        string `json:"id"`
	FullName  string `json:"full_name"`
	AppliedAt string `json:"applied_at"`
}

type SelectApplicationDto struct {
	UserId uint64 `json:"user_id"`
	JobId  uint64 `json:"job_id"`
}
