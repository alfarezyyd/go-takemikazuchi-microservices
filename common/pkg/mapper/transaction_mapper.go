package mapper

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/transaction"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func ConstructTransactionModel(jobApplicationModel *model.JobApplication, jobModel *model.Job, transactionModel *model.Transaction) {
	transactionModel.JobID = jobModel.ID
	transactionModel.PayerID = jobModel.UserId
	transactionModel.PayeeID = jobApplicationModel.ApplicantId
	transactionModel.Amount = jobModel.Price
}

func MapTransactionNotificationDtoIntoTransactionNotificationGrpc(transactionNotificationDto *dto.TransactionNotificationDto) *transaction.PostPaymentRequest {
	var transactionPostPaymentRequest transaction.PostPaymentRequest
	err := mapstructure.Decode(transactionNotificationDto, &transactionPostPaymentRequest)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &transactionPostPaymentRequest
}
