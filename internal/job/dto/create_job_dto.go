package dto

type CreateJobDto struct {
	CategoryId                   uint64  `json:"category_id"`
	Title                        string  `json:"title"`
	Description                  string  `json:"description"`
	Latitude                     float64 `json:"latitude"`
	Longitude                    float64 `json:"longitude"`
	AdditionalInformationAddress string  `json:"additional_information_address"`
	Price                        float64 `json:"price"`
}
