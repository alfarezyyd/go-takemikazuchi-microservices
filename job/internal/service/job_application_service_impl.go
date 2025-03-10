package service

import (
	"context"
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type JobApplicationServiceImpl struct {
	validatorService         validatorFeature.Service
	jobApplicationRepository repository.JobApplicationRepository
	dbConnection             *gorm.DB
	jobRepository            repository.JobRepository
	serviceDiscovery         discovery.ServiceRegistry
}

func NewJobApplicationService(
	validationInstance *validator.Validate,
	engTranslator ut.Translator,
	jobApplicationRepository repository.JobApplicationRepository,
	dbConnection *gorm.DB,
	jobRepository repository.JobRepository,
	validatorService validatorFeature.Service,
	serviceDiscovery discovery.ServiceRegistry,

) *JobApplicationServiceImpl {
	return &JobApplicationServiceImpl{
		validatorService:         validatorService,
		jobApplicationRepository: jobApplicationRepository,
		dbConnection:             dbConnection,
		jobRepository:            jobRepository,
		serviceDiscovery:         serviceDiscovery,
	}
}

func (jobApplicationService *JobApplicationServiceImpl) FindAllApplication(userJwtClaims *userDto.JwtClaimDto, jobId string) []*dto.JobApplicationResponseDto {
	err := jobApplicationService.validatorService.ValidateVar(jobId, "required,number,gt=0")
	jobApplicationService.validatorService.ParseValidationError(err)
	parsedJobId, err := strconv.ParseUint(jobId, 10, 64)
	var jobApplicationsResponse []*dto.JobApplicationResponseDto
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	err = jobApplicationService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		_, err := jobApplicationService.jobRepository.VerifyJobOwner(gormTransaction, userJwtClaims.Email, &parsedJobId)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		jobApplications := jobApplicationService.jobApplicationRepository.FindAllApplication(gormTransaction, &parsedJobId)
		jobApplicationsResponse = mapper.MapJobApplicationModelIntoJobApplicationResponse(jobApplications)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return jobApplicationsResponse
}

func (jobApplicationService *JobApplicationServiceImpl) HandleApply(ctx context.Context, userJwtClaims *userDto.JwtClaimDto, applyJobApplicationDto *dto.ApplyJobApplicationDto) {
	err := jobApplicationService.validatorService.ValidateStruct(applyJobApplicationDto)
	jobApplicationService.validatorService.ParseValidationError(err)
	err = jobApplicationService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var jobApplicationModel model.JobApplication
		userGrpcConnection, err := discovery.ServiceConnection(ctx, "userService", jobApplicationService.serviceDiscovery)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
		userAddressGrpcClient := user.NewUserServiceClient(userGrpcConnection)
		userModel, err := userAddressGrpcClient.FindByIdentifier(ctx, &user.UserIdentifier{
			Email:       helper.SafeDereference(userJwtClaims.Email, ""),
			PhoneNumber: helper.SafeDereference(userJwtClaims.PhoneNumber, ""),
		})
		exception.ParseGrpcError(err)
		isJobExists := jobApplicationService.jobRepository.IsExists(applyJobApplicationDto.JobId, gormTransaction)
		if !isJobExists {
			exception.ThrowClientError(exception.NewClientError(http.StatusNotFound, exception.ErrNotFound, errors.New("job not exists")))
		}
		jobApplicationModel.ApplicantId = userModel.ID
		jobApplicationModel.JobId = applyJobApplicationDto.JobId
		err = jobApplicationService.dbConnection.Create(&jobApplicationModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobApplicationService *JobApplicationServiceImpl) SelectApplication(ctx context.Context, userJwtClaims *userDto.JwtClaimDto, selectApplicationDto *dto.SelectApplicationDto) {
	err := jobApplicationService.validatorService.ValidateStruct(selectApplicationDto)
	jobApplicationService.validatorService.ParseValidationError(err)
	err = jobApplicationService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		jobApplicationModel := jobApplicationService.jobApplicationRepository.FindById(gormTransaction, &selectApplicationDto.UserId, &selectApplicationDto.JobId)
		userGrpcConnection, err := discovery.ServiceConnection(ctx, "userService", jobApplicationService.serviceDiscovery)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
		userAddressGrpcClient := user.NewUserServiceClient(userGrpcConnection)
		identifier, err := userAddressGrpcClient.FindByIdentifier(ctx, &user.UserIdentifier{
			Email:       helper.SafeDereference(userJwtClaims.Email, ""),
			PhoneNumber: helper.SafeDereference(userJwtClaims.PhoneNumber, ""),
		})
		exception.ParseGrpcError(err)
		id, err := jobApplicationService.jobRepository.FindVerifyById(gormTransaction, &identifier.ID, &selectApplicationDto.JobId)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		jobModel := id
		jobModel.Status = "Process"
		jobApplicationModel.Status = "Accepted"
		jobApplicationService.jobApplicationRepository.BulkRejectUpdate(gormTransaction, &jobModel.ID)
		jobApplicationService.jobApplicationRepository.Update(gormTransaction, jobApplicationModel)
		jobApplicationService.jobRepository.Update(jobModel, gormTransaction)
		return nil
	})
}
