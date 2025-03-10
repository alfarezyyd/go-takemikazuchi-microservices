package dto

import "time"

type CreateUserDto struct {
	Name            string `json:"name" validate:"required,min=3,max=100"`
	Email           string `json:"email" validate:"conditionalRequired=PhoneNumber,omitempty,email"`
	PhoneNumber     string `json:"phone_number" validate:"conditionalRequired=Email,omitempty,phoneNumber"`
	Password        string `json:"password" validate:"required,min=6,weakPassword"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,eqfield=Password"`
}

type GenerateOtpDto struct {
	UserId uint64 `json:"user_id"`
	Email  string `json:"email" validate:"required,email"`
}

type JwtClaimDto struct {
	Email       *string `json:"email" mapstructure:"email"`
	PhoneNumber *string `json:"phone_number" mapstructure:"phone_number"`
}

type LoginUserDto struct {
	UserIdentifier string `json:"user_identifier" validate:"required"`
	Password       string `json:"password" validate:"required"`
}

type UserIdentifierDto struct {
	Email       string `json:"email" validate:"omitempty,conditionalRequired=PhoneNumber,email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,conditionalRequired=Email,phoneNumber"`
}

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

type VerifyOtpDto struct {
	Email                string `json:"email" validate:"required,email"`
	OneTimePasswordToken string `json:"one_time_password_token" validate:"required,len=4"`
}
