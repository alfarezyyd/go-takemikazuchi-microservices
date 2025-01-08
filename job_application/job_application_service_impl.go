package job_application

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/helper"
	"go-takemikazuchi-api/job"
	"go-takemikazuchi-api/job_application/dto"
	"go-takemikazuchi-api/model"
	userFeature "go-takemikazuchi-api/user"
	userDto "go-takemikazuchi-api/user/dto"
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
	exception.ParseValidationError(err, jobApplicationService.engTranslator)
	err = jobApplicationService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		var jobApplicationModel model.JobApplication
		jobApplicationService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		isJobExists := jobApplicationService.jobRepository.IsExists(applyJobApplicationDto.JobId, gormTransaction)
		if !isJobExists {
			exception.ThrowClientError(exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("job not exists")))
		}
		jobApplicationModel.ApplicantId = userModel.ID
		jobApplicationModel.JobId = applyJobApplicationDto.JobId
		err = jobApplicationService.dbConnection.Create(&jobApplicationModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
