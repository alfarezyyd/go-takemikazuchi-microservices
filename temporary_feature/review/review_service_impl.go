package review

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/pkg/mapper"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices-common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/review/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/user/dto"
	userRepository "github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/user/repository"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	dbConnection     *gorm.DB
	validatorService validatorFeature.Service
	reviewRepository Repository
	jobRepository    job.Repository
	userRepository   userRepository.Repository
}

func NewService(
	dbConnection *gorm.DB,
	validatorService validatorFeature.Service,
	reviewRepository Repository,
	jobRepository job.Repository,
	userRepository userRepository.Repository,
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
