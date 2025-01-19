package job_application

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-api/internal/job_application/dto"
	exception2 "go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"net/http"
)

type Handler struct {
	jobApplicationService Service
}

func NewHandler() *Handler {
	return &Handler{}
}

func (jobApplicationHandler Handler) Apply(ginContext *gin.Context) {
	var applyJobApplication *dto.ApplyJobApplicationDto
	err := ginContext.ShouldBindBodyWithJSON(&applyJobApplication)
	helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, errors.New("bad request")))
	jobApplicationHandler.jobApplicationService.HandleApply(applyJobApplication)
}
