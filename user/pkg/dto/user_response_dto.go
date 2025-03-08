package dto

type UserResponseDto struct {
	ID              uint64
	Name            string
	Email           string
	Password        string
	Role            string
	PhoneNumber     string
	ProfilePicture  string
	IsActive        bool
	CreatedAt       string
	UpdatedAt       string
	EmailVerifiedAt *string
}
