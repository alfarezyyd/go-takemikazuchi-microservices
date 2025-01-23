package transaction

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-api/internal/transaction/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
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
	transactionHandler.transactionService.Create(userJwtClaim, &createTransactionDto)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Transaction has been created", nil))
}
