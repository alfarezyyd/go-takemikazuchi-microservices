package job_application

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/helper"
	"go-takemikazuchi-api/job_application/dto"
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
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	jobApplicationHandler.jobApplicationService.HandleApply(applyJobApplication)
}
