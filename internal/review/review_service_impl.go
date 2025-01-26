package review

import (
	"go-takemikazuchi-api/internal/job"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/internal/review/dto"
	"go-takemikazuchi-api/internal/user"
	userDto "go-takemikazuchi-api/internal/user/dto"
	validatorFeature "go-takemikazuchi-api/internal/validator"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"go-takemikazuchi-api/pkg/mapper"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	dbConnection     *gorm.DB
	validatorService validatorFeature.Service
	reviewRepository Repository
	jobRepository    job.Repository
	userRepository   user.Repository
}

func NewService(
	dbConnection *gorm.DB,
	validatorService validatorFeature.Service,
	reviewRepository Repository,
	jobRepository job.Repository,
	userRepository user.Repository,
) *ServiceImpl {
	return &ServiceImpl{
		dbConnection:     dbConnection,
		reviewRepository: reviewRepository,
		jobRepository:    jobRepository,
		userRepository:   userRepository,
		validatorService: validatorService,
	}
}

func (reviewService *ServiceImpl) Create(userJwtClaims *userDto.JwtClaimDto, createReviewDto *dto.CreateReviewDto) {
	err := reviewService.validatorService.ValidateStruct(createReviewDto)
	reviewService.validatorService.ParseValidationError(err)
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
