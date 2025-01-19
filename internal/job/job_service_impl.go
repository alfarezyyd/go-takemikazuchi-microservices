package job

import (
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/internal/category"
	dto2 "go-takemikazuchi-api/internal/job/dto"
	model2 "go-takemikazuchi-api/internal/model"
	userFeature "go-takemikazuchi-api/internal/user"
	userDto "go-takemikazuchi-api/internal/user/dto"
	exception2 "go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"go-takemikazuchi-api/pkg/mapper"
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

func (jobService *ServiceImpl) HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto *dto2.CreateJobDto) *exception2.ClientError {
	fmt.Println("adwadwa")
	fmt.Println(jobService.dbConnection)
	err := jobService.validatorInstance.Struct(createJobDto)
	exception2.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var jobModel model2.Job
		var userModel model2.User
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		isCategoryExists := jobService.categoryRepository.IsCategoryExists(createJobDto.CategoryId, gormTransaction)
		if !isCategoryExists {
			exception2.ThrowClientError(exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, errors.New("category not found")))
		}
		mapper.MapJobDtoIntoJobModel(createJobDto, &jobModel)
		jobModel.UserId = userModel.ID
		jobService.jobRepository.Store(jobModel, gormTransaction)
		return nil
	})
	helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusInternalServerError, exception2.ErrInternalServerError, err))
	return nil
}

func (jobService *ServiceImpl) HandleUpdate(userJwtClaims *userDto.JwtClaimDto, jobId string, updateJobDto *dto2.UpdateJobDto) *exception2.ClientError {
	err := jobService.validatorInstance.Struct(updateJobDto)
	exception2.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var jobModel model2.Job
		var userModel model2.User
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		mapper.MapJobDtoIntoJobModel(updateJobDto, &jobModel)
		jobModel.UserId = userModel.ID
		jobService.jobRepository.Update(jobModel, gormTransaction)
		return nil
	})
	helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusInternalServerError, exception2.ErrInternalServerError, err))
	return nil
}

func (jobService *ServiceImpl) HandleDelete(userJwtClaims *userDto.JwtClaimDto, jobId string) *exception2.ClientError {
	err := jobService.validatorInstance.Var(jobId, "required|gte=1")
	exception2.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model2.User
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		jobService.jobRepository.Delete(jobId, userModel.ID, gormTransaction)
		return nil
	})
	return nil
}
