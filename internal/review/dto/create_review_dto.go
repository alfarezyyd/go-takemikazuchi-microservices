package dto

type CreateReviewDto struct {
	ReviewedId string `json:"reviewed_id"`
	JobId      uint64 `json:"job_id"`
	Role       string `json:"role"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"review_text"`
}
