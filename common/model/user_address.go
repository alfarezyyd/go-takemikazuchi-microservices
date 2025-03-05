package model

import (
	model2 "github.com/alfarezyyd/go-takemikazuchi-microservices-job/internal/model"
)

type UserAddress struct {
	ID                    uint64       `gorm:"column:id;primaryKey;autoIncrement"`
	PlaceId               string       `gorm:"column:place_id;"`
	UserId                uint64       `gorm:"column:user_id;"`
	FormattedAddress      string       `gorm:"column:formatted_address"`
	AdditionalInformation string       `gorm:"column:additional_information"`
	StreetNumber          string       `gorm:"column:street_number;"`
	Route                 string       `gorm:"column:route;"`
	Village               string       `gorm:"column:village;"`
	District              string       `gorm:"column:district;"`
	City                  string       `gorm:"column:city;"`
	Province              string       `gorm:"column:province;"`
	Country               string       `gorm:"column:country;"`
	PostalCode            string       `gorm:"column:postal_code;"`
	Latitude              float64      `gorm:"column:latitude"`
	Longitude             float64      `gorm:"column:longitude"`
	User                  User         `gorm:"foreignKey:user_id;references:id"`
	Jobs                  []model2.Job `gorm:"foreignKey:address_id;references:id"`
}
