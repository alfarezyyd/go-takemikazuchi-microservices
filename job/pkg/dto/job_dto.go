package dto

import "time"

type CreateJobDto struct {
	AddressId                    *uint64 `form:"address_id" validate:"omitempty,gt=0"`
	CategoryId                   uint64  `form:"category_id" validate:"required,gt=0"`
	Title                        string  `form:"title" validate:"required,min=3"`
	Description                  string  `form:"description" validate:"required,min=3"`
	Latitude                     float64 `form:"latitude"`
	Longitude                    float64 `form:"longitude"`
	AdditionalInformationAddress string  `form:"additional_information_address"`
	Price                        float64 `form:"price" validate:"required,gt=0"`
}

type JobResponseDto struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type UpdateJobDto struct {
	CategoryId       uint64   `form:"category_id"`
	Title            string   `form:"title"`
	Description      string   `form:"description"`
	Price            float64  `form:"price"`
	DeletedFilesName []string `form:"deleted_files_name"`
}
