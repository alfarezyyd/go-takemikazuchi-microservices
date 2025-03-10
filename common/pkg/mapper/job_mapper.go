package mapper

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	"github.com/go-viper/mapstructure/v2"
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

func MapJobApplicationModelIntoJobApplicationResponse(jobApplicationsModel []model.JobApplication) []*dto.JobApplicationResponseDto {
	var jobApplicationsResponse []*dto.JobApplicationResponseDto
	for _, jobApplicationModel := range jobApplicationsModel {
		var jobApplicationResponseDto dto.JobApplicationResponseDto
		jobApplicationResponseDto.Id = strconv.FormatUint(jobApplicationModel.ID, 10)
		jobApplicationResponseDto.FullName = jobApplicationModel.User.Name
		jobApplicationResponseDto.AppliedAt = jobApplicationModel.CreatedAt.Format(time.RFC3339)
		jobApplicationsResponse = append(jobApplicationsResponse, &jobApplicationResponseDto)
	}
	return jobApplicationsResponse
}

func MapCreateJobDtoIntoCreateJobGrpc(createJobDto *dto.CreateJobDto) *job.CreateJobRequest {
	var createJobRequest job.CreateJobRequest
	err := mapstructure.Decode(createJobDto, &createJobRequest)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &createJobRequest
}

func MapUpdateJobDtoIntoUpdateJobGrpc(updateJobDto *dto.UpdateJobDto) *job.UpdateJobRequest {
	var updateJobRequest job.UpdateJobRequest
	err := mapstructure.Decode(updateJobDto, &updateJobRequest)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &updateJobRequest
}

func MapCreateJobGrpcIntoCreateJobDto(createJobRequest *job.CreateJobRequest) *dto.CreateJobDto {
	var createJobDto dto.CreateJobDto
	err := mapstructure.Decode(createJobRequest, &createJobDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &createJobDto
}

func MapJobModelIntoJobResponseDto(jobModel *model.Job) *dto.JobResponseDto {
	var jobResponseDto dto.JobResponseDto
	err := mapstructure.Decode(jobModel, &jobResponseDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &jobResponseDto
}

func MapJobResponseIntoJobModel(jobResponse *dto.JobResponseDto) *job.JobModel {
	var jobModel job.JobModel
	err := mapstructure.Decode(jobResponse, &jobModel)
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return &jobModel
}
