package service

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	jobDto "github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"mime/multipart"
)

type JobService interface {
	HandleCreate(ctx context.Context, userJwtClaims *userDto.JwtClaimDto, createJobDto *jobDto.CreateJobDto, uploadedFiles []*multipart.FileHeader) *exception.ClientError
	//HandleUpdate(userJwtClaims *userDto.JwtClaimDto, jobId string, updateJobDto *dto.UpdateJobDto, uploadedFiles []*multipart.FileHeader)
	//HandleDelete(userJwtClaims *userDto.JwtClaimDto, jobId string) *exception.ClientError
	//HandleRequestCompleted(userJwtClaims *userDto.JwtClaimDto, jobId *string)
}
