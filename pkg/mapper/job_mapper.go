package mapper

import (
	"github.com/go-viper/mapstructure/v2"
	dto2 "go-takemikazuchi-api/internal/job/dto"
	"go-takemikazuchi-api/internal/model"
	exception2 "go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"net/http"
)

func MapJobDtoIntoJobModel[T *dto2.CreateJobDto | *dto2.UpdateJobDto](jobDto T, jobModel *model.Job) {
	err := mapstructure.Decode(jobDto, &jobModel)
	helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
}
