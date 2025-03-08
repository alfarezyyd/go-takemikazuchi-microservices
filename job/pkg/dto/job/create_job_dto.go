package job

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
