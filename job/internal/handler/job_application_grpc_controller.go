package handler

import (
	"context"
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	jobApplication "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job_application"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/service"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type JobApplicationHandler struct {
	jobApplicationService service.JobApplicationService
	jobApplication.UnimplementedJobApplicationServiceServer
}

func NewJobApplicationHandler(grpcServer *grpc.Server, jobApplicationService service.JobApplicationService,
) *JobApplicationHandler {
	jobApplicationHandler := &JobApplicationHandler{
		jobApplicationService: jobApplicationService,
	}
	jobApplication.RegisterJobApplicationServiceServer(grpcServer, jobApplicationHandler)
	return jobApplicationHandler
}

func (jobApplicationHandler *JobApplicationHandler) FindById(ctx context.Context, findIdByRequest *jobApplication.FindJobApplicationByIdRequest) (*jobApplication.JobApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}

func (jobApplicationHandler *JobApplicationHandler) FindAllApplication(ginContext *gin.Context) {
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobId := ginContext.Param("jobId")
	jobApplicationsResponseDto := jobApplicationHandler.jobApplicationService.FindAllApplication(userJwtClaim, jobId)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Data retrieve successfully", jobApplicationsResponseDto))
}

func (jobApplicationHandler *JobApplicationHandler) SelectApplication(ginContext *gin.Context) {
	var selectApplicationDto dto.SelectApplicationDto
	err := ginContext.ShouldBindBodyWithJSON(&selectApplicationDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobApplicationHandler.jobApplicationService.SelectApplication(userJwtClaim, &selectApplicationDto)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Application successfully selected", nil))
}

func (jobApplicationHandler *JobApplicationHandler) Apply(ginContext *gin.Context) {
	var applyJobApplication *dto.ApplyJobApplicationDto
	err := ginContext.ShouldBindBodyWithJSON(&applyJobApplication)
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	jobApplicationHandler.jobApplicationService.HandleApply(userJwtClaim, applyJobApplication)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Successfully apply to the job", nil))
}
