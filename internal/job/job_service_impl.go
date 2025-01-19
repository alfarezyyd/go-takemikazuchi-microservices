package job

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/internal/category"
	jobDto "go-takemikazuchi-api/internal/job/dto"
	jobResourceFeature "go-takemikazuchi-api/internal/job_resource"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/internal/storage"
	userFeature "go-takemikazuchi-api/internal/user"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"go-takemikazuchi-api/pkg/mapper"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
)

type ServiceImpl struct {
	validatorInstance     *validator.Validate
	jobRepository         Repository
	userRepository        userFeature.Repository
	categoryRepository    category.Repository
	dbConnection          *gorm.DB
	engTranslator         ut.Translator
	jobResourceRepository jobResourceFeature.Repository
	fileStorage           storage.FileStorage
}

func NewService(validatorInstance *validator.Validate,
	jobRepository Repository,
	userRepository userFeature.Repository,
	categoryRepository category.Repository,
	jobResourceRepository jobResourceFeature.Repository,
	dbConnection *gorm.DB,
	engTranslator ut.Translator,
	fileStorage storage.FileStorage) *ServiceImpl {
	return &ServiceImpl{
		validatorInstance:     validatorInstance,
		jobRepository:         jobRepository,
		userRepository:        userRepository,
		categoryRepository:    categoryRepository,
		dbConnection:          dbConnection,
		engTranslator:         engTranslator,
		jobResourceRepository: jobResourceRepository,
		fileStorage:           fileStorage,
	}
}

func (jobService *ServiceImpl) HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto *jobDto.CreateJobDto, uploadedFiles []*multipart.FileHeader) *exception.ClientError {
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

func (jobService *ServiceImpl) HandleUpdate(userJwtClaims *userDto.JwtClaimDto, jobId string, updateJobDto *jobDto.UpdateJobDto) *exception.ClientError {
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
