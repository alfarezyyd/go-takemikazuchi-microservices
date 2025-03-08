package dto

type UserIdentifierDto struct {
	Email       string `json:"email" validate:"conditionalRequired=PhoneNumber,omitempty,email"`
	PhoneNumber string `json:"phone_number" validate:"conditionalRequired=Email,omitempty,phoneNumber"`
}
