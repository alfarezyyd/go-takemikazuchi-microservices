package dto

type UserIdentifierDto struct {
	Email       string `json:"email" validate:"omitempty,conditionalRequired=PhoneNumber,email"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,conditionalRequired=Email,phoneNumber"`
}
