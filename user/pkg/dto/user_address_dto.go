package dto

type CreateUserAddressDto struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	UserId    uint64  `json:"user_id" validate:"required"`
}

type SearchUserAddressDto struct {
	UserId        uint64 `json:"user_id" validate:"required"`
	UserAddressId uint64 `json:"user_address_id" validate:"required"`
}

type UserAddress struct {
	Industrial  string `json:"industrial"`
	Suburb      string `json:"suburb"`
	City        string `json:"city"`
	Regency     string `json:"regency"`
	State       string `json:"state"`
	ISO4        string `json:"ISO3166-2-lvl4"`
	Region      string `json:"region"`
	ISO3        string `json:"ISO3166-2-lvl3"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

// Location struct untuk merepresentasikan keseluruhan JSON
type UserLocation struct {
	PlaceID     int         `json:"place_id"`
	Licence     string      `json:"licence"`
	OsmType     string      `json:"osm_type"`
	OsmID       int         `json:"osm_id"`
	Lat         string      `json:"lat"`
	Lon         string      `json:"lon"`
	Class       string      `json:"class"`
	Type        string      `json:"type"`
	PlaceRank   int         `json:"place_rank"`
	Importance  float64     `json:"importance"`
	AddressType string      `json:"addresstype"`
	Name        string      `json:"name"`
	DisplayName string      `json:"display_name"`
	UserAddress UserAddress `json:"address"`
	BoundingBox []string    `json:"boundingbox"`
}
