package job_application

import (
	"go-takemikazuchi-api/internal/job_application/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
)

type Service interface {
	HandleApply(userJwtClaims *userDto.JwtClaimDto, applyJobApplicationDto *dto.ApplyJobApplicationDto)
}
