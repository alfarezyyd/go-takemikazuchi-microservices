package handler

import (
	"context"
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	jobApplication "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job_application"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JobApplicationHandler struct {
	serviceDiscovery discovery.ServiceRegistry
}

func NewJobApplicationHandler(
	serviceDiscovery discovery.ServiceRegistry,
) *JobApplicationHandler {
	return &JobApplicationHandler{
		serviceDiscovery: serviceDiscovery,
	}
}

func (jobApplicationHandler *JobApplicationHandler) FindAllApplication(ginContext *gin.Context) {
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobId := ginContext.Param("jobId")
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "jobService", jobApplicationHandler.serviceDiscovery)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, errors.New("job service not found")))
	jobApplicationClient := jobApplication.NewJobApplicationServiceClient(grpcConnection)
	jobApplicationsResponseDto, err := jobApplicationClient.FindAllApplication(timeoutCtx, &jobApplication.FindAllApplicationRequest{
		UserJwtClaim: mapper.MapUserJwtClaimIntoUserJwtClaimGrpc(userJwtClaim),
		JobId:        jobId,
	})
	exception.ParseGrpcError(err)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Data retrieve successfully", jobApplicationsResponseDto))
}

func (jobApplicationHandler *JobApplicationHandler) SelectApplication(ginContext *gin.Context) {
	var selectApplicationDto dto.SelectApplicationDto
	err := ginContext.ShouldBindBodyWithJSON(&selectApplicationDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "jobService", jobApplicationHandler.serviceDiscovery)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, errors.New("job service not found")))
	jobApplicationClient := jobApplication.NewJobApplicationServiceClient(grpcConnection)
	_, err = jobApplicationClient.SelectApplication(timeoutCtx, &jobApplication.SelectApplicationRequest{
		UserJwtClaim: mapper.MapUserJwtClaimIntoUserJwtClaimGrpc(userJwtClaim),
		UserId:       selectApplicationDto.UserId,
		JobId:        selectApplicationDto.JobId,
	})
	exception.ParseGrpcError(err)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Application successfully selected", nil))
}

func (jobApplicationHandler *JobApplicationHandler) Apply(ginContext *gin.Context) {
	var applyJobApplication *dto.ApplyJobApplicationDto
	err := ginContext.ShouldBindBodyWithJSON(&applyJobApplication)
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "jobService", jobApplicationHandler.serviceDiscovery)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, errors.New("job service not found")))
	jobApplicationClient := jobApplication.NewJobApplicationServiceClient(grpcConnection)
	_, err = jobApplicationClient.HandleApply(timeoutCtx, &jobApplication.ApplyRequest{
		UserJwtClaim: mapper.MapUserJwtClaimIntoUserJwtClaimGrpc(userJwtClaim),
		JobId:        applyJobApplication.JobId,
	})
	exception.ParseGrpcError(err)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Successfully apply to the job", nil))
}
