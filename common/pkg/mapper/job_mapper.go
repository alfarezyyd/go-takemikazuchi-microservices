package mapper

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	jobDto "github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto/job"
	jobApplicationDto "github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto/job_application"
	"github.com/go-viper/mapstructure/v2"
	"net/http"
	"strconv"
	"time"
)

func MapJobDtoIntoJobModel[T *jobDto.CreateJobDto | *jobDto.UpdateJobDto](jobDto T, jobModel *model.Job) {
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

func MapCreateJobDtoIntoCreateJobGrpc(createJobDto *jobDto.CreateJobDto) *job.CreateJobRequest {
	var createJobRequest job.CreateJobRequest
	err := mapstructure.Decode(createJobDto, &createJobRequest)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &createJobRequest
}

func MapUpdateJobDtoIntoUpdateJobGrpc(updateJobDto *jobDto.UpdateJobDto) *job.UpdateJobRequest {
	var updateJobRequest job.UpdateJobRequest
	err := mapstructure.Decode(updateJobDto, &updateJobRequest)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &updateJobRequest
}
