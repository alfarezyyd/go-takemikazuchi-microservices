package dto

type JobApplicationResponseDto struct {
	Id        string `json:"id"`
	FullName  string `json:"full_name"`
	AppliedAt string `json:"applied_at"`
}
