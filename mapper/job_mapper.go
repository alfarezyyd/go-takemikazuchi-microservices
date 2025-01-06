package mapper

import (
	"github.com/go-viper/mapstructure/v2"
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/helper"
	"go-takemikazuchi-api/job/dto"
	"go-takemikazuchi-api/model"
	"net/http"
)

func MapJobDtoIntoJobModel[T *dto.CreateJobDto | *dto.UpdateJobDto](jobDto T, jobModel *model.Job) {
	err := mapstructure.Decode(jobDto, &jobModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest))
}
