package dto

import "time"

type UserResponseDto struct {
	ID              uint64
	Name            string
	Email           string
	Role            string
	PhoneNumber     string
	ProfilePicture  string
	IsActive        bool
	CreatedAt       time.Time `mapstructure:"-"`
	UpdatedAt       time.Time `mapstructure:"-"`
	EmailVerifiedAt *time.Time
}
