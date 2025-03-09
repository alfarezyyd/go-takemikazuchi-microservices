package dto

type CreateUserAddressDto struct {
	Latitude  string `json:"latitude" validate:"required"`
	Longitude string `json:"longitude" validate:"required"`
	UserId    uint64 `json:"user_id" validate:"required"`
}

type SearchUserAddressDto struct {
	UserId        uint64 `json:"user_id" validate:"required"`
	UserAddressId uint64 `json:"user_address_id" validate:"required"`
}
