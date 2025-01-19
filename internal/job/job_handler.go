package job

import (
	"github.com/gin-gonic/gin"
	jobDto "go-takemikazuchi-api/internal/job/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
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
	var createJobDto jobDto.CreateJobDto
	err := ginContext.ShouldBindBodyWithJSON(&createJobDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	multipartForm, err := ginContext.MultipartForm()
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	uploadedFiles := multipartForm.File["images[]"]
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	operationResult := jobHandler.jobService.HandleCreate(userJwtClaim, &createJobDto, uploadedFiles)
	helper.CheckErrorOperation(operationResult.GetRawError(), operationResult)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Success", nil))
}

func (jobHandler *Handler) Update(ginContext *gin.Context) {
	var updateJobDto jobDto.UpdateJobDto
	err := ginContext.ShouldBindBodyWithJSON(&updateJobDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobId := ginContext.Param("id")
	operationResult := jobHandler.jobService.HandleUpdate(userJwtClaim, jobId, &updateJobDto)
	helper.CheckErrorOperation(operationResult, operationResult)
	ginContext.JSON(http.StatusOK, operationResult)
}

func (jobHandler *Handler) Delete(ginContext *gin.Context) {
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	jobId := ginContext.Param("id")
	operationResult := jobHandler.jobService.HandleDelete(userJwtClaim, jobId)
	helper.CheckErrorOperation(operationResult, operationResult)
	ginContext.JSON(http.StatusOK, operationResult)
}
