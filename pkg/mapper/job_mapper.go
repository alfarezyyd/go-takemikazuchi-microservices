package mapper

import (
	"github.com/go-viper/mapstructure/v2"
	"go-takemikazuchi-api/internal/job/dto"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"net/http"
)

func MapJobDtoIntoJobModel[T *dto.CreateJobDto | *dto.UpdateJobDto](jobDto T, jobModel *model.Job) {
	err := mapstructure.Decode(jobDto, &jobModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
}

func MapStringIntoJobResourceModel(jobId uint64, allFilePath []string) []*model.JobResource {
	var jobResourcesModel []*model.JobResource
	for _, filePath := range allFilePath {
		var jobResourceModel model.JobResource
		jobResourceModel.JobId = jobId
		jobResourceModel.ImagePath = filePath
		jobResourcesModel = append(jobResourcesModel, &jobResourceModel)
	}
	return jobResourcesModel
}
