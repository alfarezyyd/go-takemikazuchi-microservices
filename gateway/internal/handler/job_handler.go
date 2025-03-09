package handler

import (
	"context"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	jobDto "github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"time"
)

type JobHandler struct {
	serviceDiscovery discovery.ServiceRegistry
}

func NewJobHandler(serviceDiscovery discovery.ServiceRegistry) *JobHandler {
	return &JobHandler{
		serviceDiscovery: serviceDiscovery,
	}
}

func (jobHandler *JobHandler) Create(ginContext *gin.Context) {
	var createJobDto jobDto.CreateJobDto
	err := ginContext.ShouldBind(&createJobDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	var uploadedFiles []*multipart.FileHeader
	if ginContext.ContentType() == "multipart/form-data" {
		multipartForm, err := ginContext.MultipartForm()
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
		// Ambil file jika ada
		uploadedFiles = multipartForm.File["images[]"]
	}
	fmt.Println(uploadedFiles)
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "jobService", jobHandler.serviceDiscovery)
	jobClient := job.NewJobServiceClient(grpcConnection)
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	createJobGrpc := mapper.MapCreateJobDtoIntoCreateJobGrpc(&createJobDto)
	createJobGrpc.UserJwtClaim = mapper.MapUserJwtClaimIntoUserJwtClaimGrpc(userJwtClaim)
	_, clientError := jobClient.HandleCreate(timeoutCtx, createJobGrpc)
	exception.ParseGrpcError(ginContext, clientError)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Success", nil))
}

func (jobHandler *JobHandler) Update(ginContext *gin.Context) {
	var updateJobDto jobDto.UpdateJobDto
	err := ginContext.ShouldBind(&updateJobDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	var uploadedFiles []*multipart.FileHeader
	multipartForm, err := ginContext.MultipartForm()
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	uploadedFiles = multipartForm.File["images[]"]
	fmt.Println(uploadedFiles)
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobId := ginContext.Param("jobId")
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "jobService", jobHandler.serviceDiscovery)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	jobClient := job.NewJobServiceClient(grpcConnection)
	createJobGrpc := mapper.MapUpdateJobDtoIntoUpdateJobGrpc(&updateJobDto)
	createJobGrpc.UserJwtClaim = mapper.MapUserJwtClaimIntoUserJwtClaimGrpc(userJwtClaim)
	createJobGrpc.JobId = jobId
	_, clientError := jobClient.HandleUpdate(timeoutCtx, createJobGrpc)
	exception.ParseGrpcError(ginContext, clientError)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Success", nil))
}

func (jobHandler *JobHandler) Delete(ginContext *gin.Context) {
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobId := ginContext.Param("id")
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "jobService", jobHandler.serviceDiscovery)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	jobClient := job.NewJobServiceClient(grpcConnection)
	_, clientError := jobClient.HandleDelete(timeoutCtx, &job.DeleteJobRequest{
		UserJwtClaim: mapper.MapUserJwtClaimIntoUserJwtClaimGrpc(userJwtClaim),
		JobId:        jobId,
	})
	exception.ParseGrpcError(ginContext, clientError)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Success", nil))
}

func (jobHandler *JobHandler) RequestCompleted(ginContext *gin.Context) {
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobId := ginContext.Param("jobId")
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "jobService", jobHandler.serviceDiscovery)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	jobClient := job.NewJobServiceClient(grpcConnection)
	jobClient.HandleRequestCompleted(timeoutCtx, &job.JobCompleteRequest{
		UserJwtClaim: mapper.MapUserJwtClaimIntoUserJwtClaimGrpc(userJwtClaim),
		JobId:        jobId,
	})
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Success", nil))
}
