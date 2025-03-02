package review

import (
	"go-takemikazuchi-microservices/internal/review/dto"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
)

type Service interface {
	Create(userJwtClaims *userDto.JwtClaimDto, createReviewDto *dto.CreateReviewDto)
}
