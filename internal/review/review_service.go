package review

import (
	"go-takemikazuchi-api/internal/review/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
)

type Service interface {
	Create(userJwtClaims *userDto.JwtClaimDto, createReviewDto *dto.CreateReviewDto)
}
