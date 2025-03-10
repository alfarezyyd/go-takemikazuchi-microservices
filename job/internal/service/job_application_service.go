package service

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
)

type JobApplicationService interface {
	FindAllApplication(userJwtClaims *userDto.JwtClaimDto, jobId string) []*dto.JobApplicationResponseDto
	HandleApply(userJwtClaims *userDto.JwtClaimDto, applyJobApplicationDto *dto.ApplyJobApplicationDto)
	SelectApplication(userJwtClaims *userDto.JwtClaimDto, selectApplicationDto *dto.SelectApplicationDto)
}
