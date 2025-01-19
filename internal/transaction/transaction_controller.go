package transaction

import (
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-api/internal/transaction/dto"
)

type Controller interface {
	Create(ginContext *gin.Context, createTransactionDto *dto.CreateTransactionDto)
}
