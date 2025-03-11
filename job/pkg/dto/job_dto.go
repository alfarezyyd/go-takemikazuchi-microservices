package dto

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
	ID          uint64  `json:"id"`
	Title       string  `json:"title"`
	CategoryId  uint64  `json:"category_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at" mapstructure:"-"`
	UpdatedAt   string  `json:"updated_at" mapstructure:"-"`
}

type UpdateJobDto struct {
	CategoryId       uint64   `form:"category_id"`
	Title            string   `form:"title"`
	Description      string   `form:"description"`
	Price            float64  `form:"price"`
	DeletedFilesName []string `form:"deleted_files_name"`
}
