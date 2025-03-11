package service

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
)

type JobApplicationService interface {
	FindAllApplication(userJwtClaims *userDto.JwtClaimDto, jobId string) []*dto.JobApplicationResponseDto
	HandleApply(ctx context.Context, userJwtClaims *userDto.JwtClaimDto, applyJobApplicationDto *dto.ApplyJobApplicationDto)
	SelectApplication(ctx context.Context, userJwtClaims *userDto.JwtClaimDto, selectApplicationDto *dto.SelectApplicationDto)
	FindById(ctx context.Context, applicantId, jobId uint64) *dto.JobApplicationResponseDto
}
