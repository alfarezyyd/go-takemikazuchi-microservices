package job

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/helper"
	"go-takemikazuchi-api/job/dto"
	"go-takemikazuchi-api/mapper"
	"go-takemikazuchi-api/model"
	userFeature "go-takemikazuchi-api/user"
	userDto "go-takemikazuchi-api/user/dto"
	"gorm.io/gorm"
	"net/http"
)

type ServiceImpl struct {
	validatorInstance *validator.Validate
	jobRepository     Repository
	userRepository    userFeature.Repository
	dbConnection      *gorm.DB
	engTranslator     ut.Translator
}

func NewService() *ServiceImpl {
	return &ServiceImpl{}
}

func (jobService *ServiceImpl) HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto *dto.CreateJobDto) *exception.ClientError {
	err := jobService.validatorInstance.Struct(createJobDto)
	exception.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var jobModel model.Job
		var userModel model.User
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		mapper.MapJobDtoIntoJobModel(createJobDto, &jobModel)
		jobModel.UserId = userModel.ID
		jobService.jobRepository.Store(jobModel, gormTransaction)
		return nil
	})
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError))
	return nil
}

func (jobService *ServiceImpl) HandleUpdate(userJwtClaims *userDto.JwtClaimDto, updateJobDto *dto.UpdateJobDto) *exception.ClientError {
	err := jobService.validatorInstance.Struct(updateJobDto)
	exception.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var jobModel model.Job
		var userModel model.User
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		mapper.MapJobDtoIntoJobModel(updateJobDto, &jobModel)
		jobModel.UserId = userModel.ID
		jobService.jobRepository.Update(jobModel, gormTransaction)
		return nil
	})
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError))
	return nil
}
