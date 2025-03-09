package service

import (
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/repository"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type JobApplicationServiceImpl struct {
	validatorService         validatorFeature.Service
	jobApplicationRepository repository.Repository
	dbConnection             *gorm.DB
	jobRepository            repository.JobRepositoryImpl
}

func NewService(
	validationInstance *validator.Validate,
	engTranslator ut.Translator,
	jobApplicationRepository repository.Repository,
	dbConnection *gorm.DB,
	jobRepository repository.JobRepositoryImpl,
	validatorService validatorFeature.Service,
) *JobApplicationServiceImpl {
	return &JobApplicationServiceImpl{
		validatorService:         validatorService,
		jobApplicationRepository: jobApplicationRepository,
		dbConnection:             dbConnection,
		jobRepository:            jobRepository,
	}
}

//func (jobApplicationService *JobApplicationServiceImpl) FindAllApplication(userJwtClaims *userDto.JwtClaimDto, jobId string) []*dto.JobApplicationResponseDto {
//	err := jobApplicationService.validatorService.ValidateVar(jobId, "required,number,gt=0")
//	jobApplicationService.validatorService.ParseValidationError(err)
//	parsedJobId, err := strconv.ParseUint(jobId, 10, 64)
//	var jobApplicationsResponse []*dto.JobApplicationResponseDto
//	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
//	err = jobApplicationService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
//		_, err := jobApplicationService.jobRepository.VerifyJobOwner(gormTransaction, userJwtClaims.Email, &parsedJobId)
//		helper.CheckErrorOperation(err, exception.ParseGormError(err))
//		jobApplications := jobApplicationService.jobApplicationRepository.FindAllApplication(gormTransaction, &parsedJobId)
//		jobApplicationsResponse = mapper.MapJobApplicationModelIntoJobApplicationResponse(jobApplications)
//		return nil
//	})
//	helper.CheckErrorOperation(err, exception.ParseGormError(err))
//	return jobApplicationsResponse
//}
//
//func (jobApplicationService *JobApplicationServiceImpl) HandleApply(userJwtClaims *userDto.JwtClaimDto, applyJobApplicationDto *dto.ApplyJobApplicationDto) {
//	err := jobApplicationService.validatorService.ValidateStruct(applyJobApplicationDto)
//	jobApplicationService.validatorService.ParseValidationError(err)
//	err = jobApplicationService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
//		var userModel model.User
//		var jobApplicationModel model.JobApplication
//		jobApplicationService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
//		isJobExists := jobApplicationService.jobRepository.IsExists(applyJobApplicationDto.JobId, gormTransaction)
//		if !isJobExists {
//			exception.ThrowClientError(exception.NewClientError(http.StatusNotFound, exception.ErrNotFound, errors.New("job not exists")))
//		}
//		jobApplicationModel.ApplicantId = userModel.ID
//		jobApplicationModel.JobId = applyJobApplicationDto.JobId
//		err = jobApplicationService.dbConnection.Create(&jobApplicationModel).Error
//		helper.CheckErrorOperation(err, exception.ParseGormError(err))
//		return nil
//	})
//	helper.CheckErrorOperation(err, exception.ParseGormError(err))
//}
//
//func (jobApplicationService *JobApplicationServiceImpl) SelectApplication(userJwtClaims *userDto.JwtClaimDto, selectApplicationDto *dto.SelectApplicationDto) {
//	err := jobApplicationService.validatorService.ValidateStruct(selectApplicationDto)
//	jobApplicationService.validatorService.ParseValidationError(err)
//	err = jobApplicationService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
//		jobApplicationModel := jobApplicationService.jobApplicationRepository.FindById(gormTransaction, &selectApplicationDto.UserId, &selectApplicationDto.JobId)
//		id, err := jobApplicationService.jobRepository.FindVerifyById(gormTransaction, userJwtClaims.Email, &selectApplicationDto.JobId)
//		helper.CheckErrorOperation(err, exception.ParseGormError(err))
//		jobModel := id
//		jobModel.Status = "Process"
//		jobApplicationModel.Status = "Accepted"
//		jobApplicationService.jobApplicationRepository.BulkRejectUpdate(gormTransaction, &jobModel.ID)
//		jobApplicationService.jobApplicationRepository.Update(gormTransaction, jobApplicationModel)
//		jobApplicationService.jobRepository.Update(jobModel, gormTransaction)
//		return nil
//	})
//}
