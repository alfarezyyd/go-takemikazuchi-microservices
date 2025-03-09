package service

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/internal/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type TransactionServiceImpl struct {
	validatorInstance        *validator.Validate
	engTranslator            ut.Translator
	gormTransaction          *gorm.DB
	midtransClient           *snap.Client
	jobRepository            job.Repository
	jobApplicationRepository jobApplication.Repository
	transactionRepository    repository.Repository
	viperConfig              *viper.Viper
}

func NewTransactionService(
	validatorInstance *validator.Validate,
	engTranslator ut.Translator,
	gormTransaction *gorm.DB,
	midtransClient *snap.Client,
	jobRepository job.Repository,
	transactionRepository repository.Repository,
	jobApplicationRepository jobApplication.Repository,
	viperConfig *viper.Viper,
) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		validatorInstance:        validatorInstance,
		engTranslator:            engTranslator,
		gormTransaction:          gormTransaction,
		midtransClient:           midtransClient,
		jobRepository:            jobRepository,
		transactionRepository:    transactionRepository,
		viperConfig:              viperConfig,
		jobApplicationRepository: jobApplicationRepository,
	}
}

func (transactionService *TransactionServiceImpl) Create(userJwtClaims *userDto.JwtClaimDto, createTransactionDto *dto.CreateTransactionDto) string {
	var midtransSnapToken string
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
		transactionModel.SnapToken = &midtransResponse.Token
		midtransSnapToken = midtransResponse.Token
		transactionService.transactionRepository.Update(gormTransaction, &transactionModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return midtransSnapToken
}

func (transactionService *TransactionServiceImpl) PostPayment(transactionNotificationDto *dto.TransactionNotificationDto) {
	err := transactionService.gormTransaction.Transaction(func(gormTransaction *gorm.DB) error {
		transactionModel := transactionService.transactionRepository.FindWithRelationship(gormTransaction, transactionNotificationDto.OrderId)
		//SHA512(order_id+status_code+gross_amount+ServerKey)
		notificationSignatureKey := fmt.Sprintf("%s%s%.2f%s", transactionModel.ID, transactionNotificationDto.StatusCode, transactionModel.Amount, transactionService.viperConfig.GetString("MIDTRANS_SERVER_KEY"))
		generatedHash := sha512.Sum512([]byte(notificationSignatureKey))
		generatedHexadecimalHash := hex.EncodeToString(generatedHash[:])
		if generatedHexadecimalHash != transactionNotificationDto.SignatureKey {
			exception.ThrowClientError(exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("error when trying to hash")))
		}
		transactionService.paymentOperation(transactionNotificationDto, transactionModel)
		transactionModel.PaymentMethod = &transactionNotificationDto.PaymentType
		transactionService.transactionRepository.Update(gormTransaction, transactionModel)
		transactionService.jobRepository.Update(transactionModel.Job, gormTransaction)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (transactionService *TransactionServiceImpl) paymentOperation(transactionNotificationDto *dto.TransactionNotificationDto, transactionModel *model.Transaction) {
	if transactionNotificationDto != nil {
		switch transactionNotificationDto.TransactionStatus {
		case "capture":
			if transactionNotificationDto.FraudStatus == "challenge" {
				// TODO set transaction status on your database to 'challenge'
				// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			} else if transactionNotificationDto.FraudStatus == "accept" {
				transactionModel.Job.Status = "On Working"
				transactionModel.Status = "Completed"
			}
			break
		case "settlement":
			transactionModel.Job.Status = "On Working"
			transactionModel.Status = "Completed"
			break
		case "deny":
			// TODO you can ignore 'deny', because most of the time it allows payment retries
			// and later can become success
			break
		case "cancel":
		case "expire":
			transactionModel.Job.Status = "Closed"
			transactionModel.Status = "Failed"
			// TODO set transaction status on your databaase to 'failure'
			break
		case "pending":
			break
		}
	}
}
