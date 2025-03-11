package service

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/category"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job"
	jobApplication "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job_application"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
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
	"log"
	"net/http"
	"strconv"
)

type TransactionServiceImpl struct {
	validatorInstance *validator.Validate
	engTranslator     ut.Translator
	gormTransaction   *gorm.DB
	midtransClient    *snap.Client
	//jobRepository            job.Repository
	//jobApplicationRepository jobApplication.Repository
	transactionRepository repository.TransactionRepository
	viperConfig           *viper.Viper
	serviceRegistry       discovery.ServiceRegistry
	rabbitMQ              *configs.RabbitMQ
}

func NewTransactionService(
	validatorInstance *validator.Validate,
	engTranslator ut.Translator,
	gormTransaction *gorm.DB,
	midtransClient *snap.Client,
	//jobRepository job.Repository,
	transactionRepository repository.TransactionRepository,
	//jobApplicationRepository jobApplication.Repository,
	viperConfig *viper.Viper,
	serviceRegistry discovery.ServiceRegistry,
	rabbitMQ *configs.RabbitMQ,
) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		validatorInstance: validatorInstance,
		engTranslator:     engTranslator,
		gormTransaction:   gormTransaction,
		midtransClient:    midtransClient,
		//jobRepository:            jobRepository,
		transactionRepository: transactionRepository,
		viperConfig:           viperConfig,
		serviceRegistry:       serviceRegistry,
		rabbitMQ:              rabbitMQ,
		//jobApplicationRepository: jobApplicationRepository,
	}
}

func (transactionService *TransactionServiceImpl) Create(ctx context.Context, userJwtClaims *userDto.JwtClaimDto, createTransactionDto *dto.CreateTransactionDto) string {
	var midtransSnapToken string
	err := transactionService.validatorInstance.Struct(createTransactionDto)
	exception.ParseValidationError(err, transactionService.engTranslator)
	err = transactionService.gormTransaction.Transaction(func(gormTransaction *gorm.DB) error {
		grpcUserConnection, err := discovery.ServiceConnection(ctx, "userService", transactionService.serviceRegistry)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, errors.New("job service not found")))
		grpcUserClient := user.NewUserServiceClient(grpcUserConnection)
		userGrpcResponse, err := grpcUserClient.FindByIdentifier(ctx, &user.UserIdentifier{
			Email:       helper.SafeDereference(userJwtClaims.Email, ""),
			PhoneNumber: helper.SafeDereference(userJwtClaims.PhoneNumber, ""),
		})
		grpcJobConnection, err := discovery.ServiceConnection(ctx, "jobService", transactionService.serviceRegistry)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, errors.New("job service not found")))
		grpcJobClient := job.NewJobServiceClient(grpcJobConnection)
		jobModel, err := grpcJobClient.FindById(ctx, &job.FindByIdRequest{
			JobId:     createTransactionDto.JobId,
			UserEmail: userGrpcResponse.Email,
		})

		grpcJobApplicationClient := jobApplication.NewJobApplicationServiceClient(grpcJobConnection)
		jobApplicationModel, err := grpcJobApplicationClient.FindById(ctx, &jobApplication.FindJobApplicationByIdRequest{
			ApplicantId: createTransactionDto.ApplicantId,
			JobId:       createTransactionDto.JobId,
		})

		grpcCategoryConnection, err := discovery.ServiceConnection(ctx, "categoryService", transactionService.serviceRegistry)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, errors.New("job service not found")))
		grpcCategoryClient := category.NewCategoryServiceClient(grpcCategoryConnection)
		categoryModel, err := grpcCategoryClient.FindById(ctx, &category.SearchCategoryRequest{CategoryId: jobModel.CategoryId})
		uuidString := fmt.Sprintf("%s-%s", "order", uuid.New().String())
		var transactionModel model.Transaction
		transactionModel.ID = uuidString
		mapper.ConstructTransactionModel(jobApplicationModel, jobModel, &transactionModel, userGrpcResponse.ID)
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
					Category: categoryModel.Name,
				},
			},
			CustomerDetail: &midtrans.CustomerDetails{
				FName: userGrpcResponse.Name,
				LName: "",
				Email: userGrpcResponse.Email,
				Phone: helper.ParseNullableValue(userGrpcResponse.PhoneNumber),
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
		//transactionService.jobRepository.Update(transactionModel.Job, gormTransaction)

		// Publish ke RabbitMQ
		orderUpdateMessage := map[string]interface{}{
			"order_id":   transactionModel.ID,
			"status":     transactionNotificationDto.TransactionStatus,
			"payment":    transactionNotificationDto.PaymentType,
			"updated_at": transactionModel.UpdatedAt,
			"job_id":     strconv.FormatUint(transactionModel.JobID, 10),
		}

		err := transactionService.rabbitMQ.Publish(orderUpdateMessage)
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
			return err
		}

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
				transactionModel.Status = "Completed"
			}
			break
		case "settlement":
			transactionModel.Status = "Completed"
			break
		case "deny":
			// TODO you can ignore 'deny', because most of the time it allows payment retries
			// and later can become success
			break
		case "cancel":
		case "expire":
			transactionModel.Status = "Failed"
			// TODO set transaction status on your databaase to 'failure'
			break
		case "pending":
			break
		}
	}
}
