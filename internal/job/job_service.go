package job

import (
	dto2 "go-takemikazuchi-api/internal/job/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
)

type Service interface {
	HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto *dto2.CreateJobDto) *exception.ClientError
	HandleUpdate(userJwtClaims *userDto.JwtClaimDto, jobId string, updateJobDto *dto2.UpdateJobDto) *exception.ClientError
	HandleDelete(userJwtClaims *userDto.JwtClaimDto, jobId string) *exception.ClientError
}
