package handler

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Create(ginContext *gin.Context)
	Notification(ginContext *gin.Context)
}
