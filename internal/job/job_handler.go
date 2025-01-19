package job

import (
	"github.com/gin-gonic/gin"
	dto2 "go-takemikazuchi-api/internal/job/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
	exception2 "go-takemikazuchi-api/pkg/exception"
	helper2 "go-takemikazuchi-api/pkg/helper"
	"net/http"
)

type Handler struct {
	jobService Service
}

func NewHandler(jobService Service) *Handler {
	return &Handler{
		jobService: jobService,
	}
}

func (jobHandler *Handler) Create(ginContext *gin.Context) {
	var createJobDto dto2.CreateJobDto
	err := ginContext.ShouldBindBodyWithJSON(&createJobDto)
	helper2.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	operationResult := jobHandler.jobService.HandleCreate(userJwtClaim, &createJobDto)
	helper2.CheckErrorOperation(operationResult.GetRawError(), operationResult)
	ginContext.JSON(http.StatusCreated, helper2.WriteSuccess("Success", nil))
}

func (jobHandler *Handler) Update(ginContext *gin.Context) {
	var updateJobDto dto2.UpdateJobDto
	err := ginContext.ShouldBindBodyWithJSON(&updateJobDto)
	helper2.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobId := ginContext.Param("id")
	operationResult := jobHandler.jobService.HandleUpdate(userJwtClaim, jobId, &updateJobDto)
	helper2.CheckErrorOperation(operationResult, operationResult)
	ginContext.JSON(http.StatusOK, operationResult)
}

func (jobHandler *Handler) Delete(ginContext *gin.Context) {
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobId := ginContext.Param("id")
	operationResult := jobHandler.jobService.HandleDelete(userJwtClaim, jobId)
	helper2.CheckErrorOperation(operationResult, operationResult)
	ginContext.JSON(http.StatusOK, operationResult)
}
