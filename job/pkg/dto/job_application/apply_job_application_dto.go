package job_application

type ApplyJobApplicationDto struct {
	JobId uint64 `json:"job_id" validate:"required,gte=1"`
}
