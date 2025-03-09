package service

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/category/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto/user"
	"mime/multipart"
)

type Service interface {
	HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto *dto.CreateJobDto, uploadedFiles []*multipart.FileHeader) *exception.ClientError
	//HandleUpdate(userJwtClaims *userDto.JwtClaimDto, jobId string, updateJobDto *dto.UpdateJobDto, uploadedFiles []*multipart.FileHeader)
	//HandleDelete(userJwtClaims *userDto.JwtClaimDto, jobId string) *exception.ClientError
	//HandleRequestCompleted(userJwtClaims *userDto.JwtClaimDto, jobId *string)
}
