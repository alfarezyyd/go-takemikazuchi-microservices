package dto

type CreateUserDto struct {
	Name            string `json:"name" validate:"required,min=3,max=100"`
	Email           string `json:"email" validate:"conditionalRequired=PhoneNumber,omitempty,email"`
	PhoneNumber     string `json:"phone_number" validate:"conditionalRequired=Email,omitempty,phoneNumber"`
	Password        string `json:"password" validate:"required,min=6,weakPassword"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6,eqfield=Password"`
}
