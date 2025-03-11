package handler

import (
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	Create(ginContext *gin.Context)
	Notification(ginContext *gin.Context)
}
