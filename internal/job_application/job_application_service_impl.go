package job_application

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/internal/job"
	"go-takemikazuchi-api/internal/job_application/dto"
	model2 "go-takemikazuchi-api/internal/model"
	userFeature "go-takemikazuchi-api/internal/user"
	userDto "go-takemikazuchi-api/internal/user/dto"
	exception2 "go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"gorm.io/gorm"
	"net/http"
)

type ServiceImpl struct {
	validationInstance       *validator.Validate
	engTranslator            ut.Translator
	jobApplicationRepository Repository
	dbConnection             *gorm.DB
	jobRepository            job.Repository
	userRepository           userFeature.Repository
}

func NewService() *ServiceImpl {
	return &ServiceImpl{}
}

func (jobApplicationService *ServiceImpl) HandleApply(userJwtClaims *userDto.JwtClaimDto, applyJobApplicationDto *dto.ApplyJobApplicationDto) {
	err := jobApplicationService.validationInstance.Struct(applyJobApplicationDto)
	exception2.ParseValidationError(err, jobApplicationService.engTranslator)
	err = jobApplicationService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model2.User
		var jobApplicationModel model2.JobApplication
		jobApplicationService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		isJobExists := jobApplicationService.jobRepository.IsExists(applyJobApplicationDto.JobId, gormTransaction)
		if !isJobExists {
			exception2.ThrowClientError(exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, errors.New("job not exists")))
		}
		jobApplicationModel.ApplicantId = userModel.ID
		jobApplicationModel.JobId = applyJobApplicationDto.JobId
		err = jobApplicationService.dbConnection.Create(&jobApplicationModel).Error
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		return nil
	})
	helper.CheckErrorOperation(err, exception2.ParseGormError(err))
}
