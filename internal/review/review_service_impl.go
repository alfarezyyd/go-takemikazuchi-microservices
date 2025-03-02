package review

import (
	"go-takemikazuchi-microservices/internal/job"
	"go-takemikazuchi-microservices/internal/model"
	"go-takemikazuchi-microservices/internal/review/dto"
	"go-takemikazuchi-microservices/internal/user"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
	validatorFeature "go-takemikazuchi-microservices/internal/validator"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"go-takemikazuchi-microservices/pkg/mapper"
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
