package review

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/internal/job"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/internal/review/dto"
	"go-takemikazuchi-api/internal/user"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"go-takemikazuchi-api/pkg/mapper"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	dbConnection      *gorm.DB
	validatorInstance *validator.Validate
	reviewRepository  Repository
	engTranslator     ut.Translator
	jobRepository     job.Repository
	userRepository    user.Repository
}

func NewService(
	dbConnection *gorm.DB,
	validatorInstance *validator.Validate,
	reviewRepository Repository,
	engTranslator ut.Translator,
	jobRepository job.Repository,
	userRepository user.Repository,
) *ServiceImpl {
	return &ServiceImpl{
		validatorInstance: validatorInstance,
		dbConnection:      dbConnection,
		reviewRepository:  reviewRepository,
		engTranslator:     engTranslator,
		jobRepository:     jobRepository,
		userRepository:    userRepository,
	}
}

func (reviewService *ServiceImpl) Create(userJwtClaims *userDto.JwtClaimDto, createReviewDto *dto.CreateReviewDto) {
	err := reviewService.validatorInstance.Struct(createReviewDto)
	exception.ParseValidationError(err, reviewService.engTranslator)
	err = reviewService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		reviewService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		_, err := reviewService.jobRepository.VerifyJobOwner(gormTransaction, &userModel.Email, &createReviewDto.JobId)
		if err != nil {
			_, err := reviewService.jobRepository.VerifyJobWorker(gormTransaction, &userModel.Email, &createReviewDto.JobId)
			helper.CheckErrorOperation(err, exception.ParseGormError(err))
		}
		var reviewModel model.Review
		mapper.MapCreateReviewDtoIntoReviewModel(createReviewDto, &reviewModel)
		reviewModel.ReviewerId = userModel.ID
		reviewService.reviewRepository.Create(gormTransaction, &reviewModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
