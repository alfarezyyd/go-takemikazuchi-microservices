package service

import (
	"go-takemikazuchi-microservices/internal/job_application/dto"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
)

type Service interface {
	FindAllApplication(userJwtClaims *userDto.JwtClaimDto, jobId string) []*dto.JobApplicationResponseDto
	HandleApply(userJwtClaims *userDto.JwtClaimDto, applyJobApplicationDto *dto.ApplyJobApplicationDto)
	SelectApplication(userJwtClaims *userDto.JwtClaimDto, selectApplicationDto *dto.SelectApplicationDto)
}
