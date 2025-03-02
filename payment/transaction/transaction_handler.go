package transaction

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-microservices/internal/transaction/dto"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"net/http"
)

type Handler struct {
	transactionService Service
}

func NewHandler(transactionService Service) *Handler {
	return &Handler{
		transactionService: transactionService,
	}
}

func (transactionHandler *Handler) Create(ginContext *gin.Context) {
	var createTransactionDto dto.CreateTransactionDto
	err := ginContext.ShouldBindBodyWithJSON(&createTransactionDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("failed to parse body")))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	midtransSnapToken := transactionHandler.transactionService.Create(userJwtClaim, &createTransactionDto)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Transaction has been created", gin.H{
		"token": midtransSnapToken,
	}))
}

func (transactionHandler *Handler) Notification(ginContext *gin.Context) {
	var transactionNotificationDto dto.TransactionNotificationDto
	err := ginContext.ShouldBindBodyWithJSON(&transactionNotificationDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody, errors.New("failed to parse body")))
	transactionHandler.transactionService.PostPayment(&transactionNotificationDto)
}
