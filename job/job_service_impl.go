package job

import (
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/category"
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
	validatorInstance  *validator.Validate
	jobRepository      Repository
	userRepository     userFeature.Repository
	categoryRepository category.Repository
	dbConnection       *gorm.DB
	engTranslator      ut.Translator
}

func NewService(validatorInstance *validator.Validate,
	jobRepository Repository,
	userRepository userFeature.Repository,
	categoryRepository category.Repository,
	dbConnection *gorm.DB,
	engTranslator ut.Translator) *ServiceImpl {
	return &ServiceImpl{
		validatorInstance:  validatorInstance,
		jobRepository:      jobRepository,
		userRepository:     userRepository,
		categoryRepository: categoryRepository,
		dbConnection:       dbConnection,
		engTranslator:      engTranslator,
	}
}

func (jobService *ServiceImpl) HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto *dto.CreateJobDto) *exception.ClientError {
	fmt.Println("adwadwa")
	fmt.Println(jobService.dbConnection)
	err := jobService.validatorInstance.Struct(createJobDto)
	exception.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var jobModel model.Job
		var userModel model.User
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		isCategoryExists := jobService.categoryRepository.IsCategoryExists(createJobDto.CategoryId, gormTransaction)
		if !isCategoryExists {
			exception.ThrowClientError(exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("category not found")))
		}
		mapper.MapJobDtoIntoJobModel(createJobDto, &jobModel)
		jobModel.UserId = userModel.ID
		jobService.jobRepository.Store(jobModel, gormTransaction)
		return nil
	})
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	return nil
}

func (jobService *ServiceImpl) HandleUpdate(userJwtClaims *userDto.JwtClaimDto, jobId string, updateJobDto *dto.UpdateJobDto) *exception.ClientError {
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
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	return nil
}

func (jobService *ServiceImpl) HandleDelete(userJwtClaims *userDto.JwtClaimDto, jobId string) *exception.ClientError {
	err := jobService.validatorInstance.Var(jobId, "required|gte=1")
	exception.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		jobService.jobRepository.Delete(jobId, userModel.ID, gormTransaction)
		return nil
	})
	return nil
}
