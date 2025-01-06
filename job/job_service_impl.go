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

func (serviceImpl *ServiceImpl) HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto *dto.CreateJobDto) *exception.ClientError {
	err := serviceImpl.validatorInstance.Struct(createJobDto)
	exception.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var jobModel model.Job
		var userModel model.User
		serviceImpl.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		mapper.MapJobDtoIntoJobModel(createJobDto, &jobModel)
		jobModel.UserId = userModel.ID
		serviceImpl.jobRepository.Store(jobModel, gormTransaction)
		return nil
	})
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError))
	return nil
}
