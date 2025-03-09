package handler

import (
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/service"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	jobApplicationService service.Service
}

func NewHandler(
	jobApplicationService service.Service,
) *Handler {
	return &Handler{
		jobApplicationService: jobApplicationService,
	}
}

func (jobApplicationHandler Handler) FindAllApplication(ginContext *gin.Context) {
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobId := ginContext.Param("jobId")
	jobApplicationsResponseDto := jobApplicationHandler.jobApplicationService.FindAllApplication(userJwtClaim, jobId)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Data retrieve successfully", jobApplicationsResponseDto))
}

func (jobApplicationHandler Handler) SelectApplication(ginContext *gin.Context) {
	var selectApplicationDto dto.SelectApplicationDto
	err := ginContext.ShouldBindBodyWithJSON(&selectApplicationDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobApplicationHandler.jobApplicationService.SelectApplication(userJwtClaim, &selectApplicationDto)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Application successfully selected", nil))
}

func (jobApplicationHandler Handler) Apply(ginContext *gin.Context) {
	var applyJobApplication *dto.ApplyJobApplicationDto
	err := ginContext.ShouldBindBodyWithJSON(&applyJobApplication)
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	jobApplicationHandler.jobApplicationService.HandleApply(userJwtClaim, applyJobApplication)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Successfully apply to the job", nil))
}
