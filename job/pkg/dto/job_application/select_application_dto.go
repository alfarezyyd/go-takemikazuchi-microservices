package job_application

type SelectApplicationDto struct {
	UserId uint64 `json:"user_id"`
	JobId  uint64 `json:"job_id"`
}
