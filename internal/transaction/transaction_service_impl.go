package transaction

import (
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"go-takemikazuchi-api/internal/job"
	jobApplication "go-takemikazuchi-api/internal/job_application"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/internal/transaction/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"go-takemikazuchi-api/pkg/mapper"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ServiceImpl struct {
	validatorInstance        *validator.Validate
	engTranslator            ut.Translator
	gormTransaction          *gorm.DB
	midtransClient           *snap.Client
	jobRepository            job.Repository
	jobApplicationRepository jobApplication.Repository
	transactionRepository    Repository
}

func NewService(
	validatorInstance *validator.Validate,
	engTranslator ut.Translator,
	gormTransaction *gorm.DB,
	midtransClient *snap.Client,
	jobRepository job.Repository,
	transactionRepository Repository,
	jobApplicationRepository jobApplication.Repository,

) *ServiceImpl {
	return &ServiceImpl{
		validatorInstance:        validatorInstance,
		engTranslator:            engTranslator,
		gormTransaction:          gormTransaction,
		midtransClient:           midtransClient,
		jobRepository:            jobRepository,
		transactionRepository:    transactionRepository,
		jobApplicationRepository: jobApplicationRepository,
	}
}

func (transactionService *ServiceImpl) Create(userJwtClaims *userDto.JwtClaimDto, createTransactionDto *dto.CreateTransactionDto) {
	err := transactionService.validatorInstance.Struct(createTransactionDto)
	exception.ParseValidationError(err, transactionService.engTranslator)
	err = transactionService.gormTransaction.Transaction(func(gormTransaction *gorm.DB) error {
		jobModel := transactionService.jobRepository.FindWithRelationship(gormTransaction, userJwtClaims.Email, &createTransactionDto.JobId)
		jobApplicationModel := transactionService.jobApplicationRepository.FindById(gormTransaction, &createTransactionDto.ApplicantId, &createTransactionDto.JobId)
		uuidString := fmt.Sprintf("%s-%s", "order", uuid.New().String())
		var transactionModel model.Transaction
		transactionModel.ID = uuidString
		mapper.ConstructTransactionModel(jobApplicationModel, jobModel, &transactionModel)
		transactionService.transactionRepository.Create(gormTransaction, &transactionModel)
		midtransResponse, midtransError := transactionService.midtransClient.CreateTransaction(&snap.Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  uuidString,
				GrossAmt: int64(jobModel.Price),
			},
			Items: &[]midtrans.ItemDetails{
				{
					ID:       strconv.FormatUint(jobModel.ID, 10),
					Name:     jobModel.Title,
					Price:    int64(jobModel.Price),
					Qty:      1,
					Category: jobModel.Category.Name,
				},
			},
			CustomerDetail: &midtrans.CustomerDetails{
				FName: jobModel.User.Name,
				LName: "",
				Email: jobModel.User.Email,
				Phone: helper.ParseNullableValue(jobModel.User.PhoneNumber),
			},
		})
		if midtransError != nil && helper.CheckErrorOperation(midtransError.GetRawError(), exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody, errors.New("error when create midtrans transaction"))) {
			return nil
		}
		transactionModel.SnapToken = midtransResponse.Token
		transactionService.transactionRepository.Update(gormTransaction, &transactionModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
