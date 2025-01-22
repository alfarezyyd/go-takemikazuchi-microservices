package dto

type CreateJobDto struct {
	AddressId                    *uint64 `json:"address_id" validate:"gt=0"`
	CategoryId                   uint64  `json:"category_id" validate:"required,gt=0"`
	Title                        string  `json:"title" validate:"required,min=3"`
	Description                  string  `json:"description" validate:"required,min=3"`
	Latitude                     float64 `json:"latitude"`
	Longitude                    float64 `json:"longitude"`
	AdditionalInformationAddress string  `json:"additional_information_address"`
	Price                        float64 `json:"price" validate:"required,gt=0"`
}
