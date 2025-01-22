package mapper

import (
	"github.com/go-viper/mapstructure/v2"
	"go-takemikazuchi-api/internal/job/dto"
	jobApplicationDto "go-takemikazuchi-api/internal/job_application/dto"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"net/http"
	"strconv"
	"time"
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

func MapJobApplicationModelIntoJobApplicationResponse(jobApplicationsModel []model.JobApplication) []*jobApplicationDto.JobApplicationResponseDto {
	var jobApplicationsResponse []*jobApplicationDto.JobApplicationResponseDto
	for _, jobApplicationModel := range jobApplicationsModel {
		var jobApplicationResponseDto jobApplicationDto.JobApplicationResponseDto
		jobApplicationResponseDto.Id = strconv.FormatUint(jobApplicationModel.ID, 10)
		jobApplicationResponseDto.FullName = jobApplicationModel.User.Name
		jobApplicationResponseDto.AppliedAt = jobApplicationModel.CreatedAt.Format(time.RFC3339)
		jobApplicationsResponse = append(jobApplicationsResponse, &jobApplicationResponseDto)
	}
	return jobApplicationsResponse
}
