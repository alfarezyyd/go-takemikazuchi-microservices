package review

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/review/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/user/dto"
)

type Service interface {
	Create(userJwtClaims *userDto.JwtClaimDto, createReviewDto *dto.CreateReviewDto)
}
