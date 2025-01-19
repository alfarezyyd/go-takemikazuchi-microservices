package dto

type UpdateJobDto struct {
	CategoryId  uint64  `json:"category_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
