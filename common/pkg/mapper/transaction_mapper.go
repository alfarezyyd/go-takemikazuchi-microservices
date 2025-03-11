package mapper

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job"
	jobApplication "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job_application"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/transaction"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func ConstructTransactionModel(jobApplicationModel *jobApplication.JobApplicationResponse, jobModel *job.JobModel, transactionModel *model.Transaction, userId uint64) {
	transactionModel.JobID = jobModel.ID
	transactionModel.PayerID = userId
	transactionModel.PayeeID = jobApplicationModel.ApplicantId
	transactionModel.Amount = jobModel.Price
}

func MapTransactionNotificationDtoIntoTransactionNotificationGrpc(transactionNotificationDto *dto.TransactionNotificationDto) *transaction.PostPaymentRequest {
	var transactionPostPaymentRequest transaction.PostPaymentRequest
	err := mapstructure.Decode(transactionNotificationDto, &transactionPostPaymentRequest)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &transactionPostPaymentRequest
}

func MapTransactionNotificationGrpcIntoTransactionNotificationDto(postPaymentRequest *transaction.PostPaymentRequest) *dto.TransactionNotificationDto {
	var transactionNotificationDto dto.TransactionNotificationDto
	err := mapstructure.Decode(postPaymentRequest, &transactionNotificationDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &transactionNotificationDto
}
