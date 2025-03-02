package job

import (
	"go-takemikazuchi-microservices/internal/job/dto"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
	"go-takemikazuchi-microservices/pkg/exception"
	"mime/multipart"
)

type Service interface {
	HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto *dto.CreateJobDto, uploadedFiles []*multipart.FileHeader) *exception.ClientError
	HandleUpdate(userJwtClaims *userDto.JwtClaimDto, jobId string, updateJobDto *dto.UpdateJobDto, uploadedFiles []*multipart.FileHeader)
	HandleDelete(userJwtClaims *userDto.JwtClaimDto, jobId string) *exception.ClientError
	HandleRequestCompleted(userJwtClaims *userDto.JwtClaimDto, jobId *string)
}
