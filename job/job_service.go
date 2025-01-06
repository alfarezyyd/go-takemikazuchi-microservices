package job

import (
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/job/dto"
	userDto "go-takemikazuchi-api/user/dto"
)

type Service interface {
	HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto dto.CreateJobDto) *exception.ClientError
	HandleUpdate(userJwtClaims *userDto.JwtClaimDto, jobId string, updateJobDto dto.UpdateJobDto) *exception.ClientError
	HandleDelete(userJwtClaims *userDto.JwtClaimDto, jobId string) *exception.ClientError
}
