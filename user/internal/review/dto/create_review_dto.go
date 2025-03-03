package dto

type CreateReviewDto struct {
	ReviewedId uint64 `json:"reviewed_id"`
	JobId      uint64 `json:"job_id"`
	Role       string `json:"role"`
	Rating     byte   `json:"rating"`
	ReviewText string `json:"review_text"`
}
