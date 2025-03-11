package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/transaction"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TransactionHandler struct {
	serviceDiscovery discovery.ServiceRegistry
}

func NewTransactionHandler(serviceDiscovery discovery.ServiceRegistry) *TransactionHandler {
	return &TransactionHandler{
		serviceDiscovery: serviceDiscovery,
	}
}

func (transactionHandler *TransactionHandler) Create(ginContext *gin.Context) {
	var createTransactionDto dto.CreateTransactionDto
	err := ginContext.ShouldBindBodyWithJSON(&createTransactionDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("failed to parse body")))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "paymentService", transactionHandler.serviceDiscovery)
	transactionClient := transaction.NewTransactionServiceClient(grpcConnection)
	midtransSnapToken, err := transactionClient.Create(timeoutCtx, &transaction.CreateTransactionRequest{
		UserJwtClaim: mapper.MapUserJwtClaimIntoUserJwtClaimGrpc(userJwtClaim),
		JobId:        createTransactionDto.JobId,
		ApplicantId:  createTransactionDto.ApplicantId,
	})
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("failed to parse body")))
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Transaction has been created", gin.H{
		"token": midtransSnapToken.SnapToken,
	}))
}

func (transactionHandler *TransactionHandler) Notification(ginContext *gin.Context) {
	var transactionNotificationDto dto.TransactionNotificationDto
	err := ginContext.ShouldBindBodyWithJSON(&transactionNotificationDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody, errors.New("failed to parse body")))
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "paymentService", transactionHandler.serviceDiscovery)
	transactionClient := transaction.NewTransactionServiceClient(grpcConnection)
	transactionGrpc := mapper.MapTransactionNotificationDtoIntoTransactionNotificationGrpc(&transactionNotificationDto)
	_, err = transactionClient.PostPayment(timeoutCtx, transactionGrpc)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("failed to parse body")))
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Transaction has been created", nil))
}
