package job_application

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-api/internal/job_application/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"net/http"
)

type Handler struct {
	jobApplicationService Service
}

func NewHandler(
	jobApplicationService Service,
) *Handler {
	return &Handler{
		jobApplicationService: jobApplicationService,
	}
}

func (jobApplicationHandler Handler) Apply(ginContext *gin.Context) {
	var applyJobApplication *dto.ApplyJobApplicationDto
	err := ginContext.ShouldBindBodyWithJSON(&applyJobApplication)
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	jobApplicationHandler.jobApplicationService.HandleApply(userJwtClaim, applyJobApplication)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("User created successfully", nil))
}
