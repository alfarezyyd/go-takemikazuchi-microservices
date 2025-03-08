package handler

import "github.com/gin-gonic/gin"

type JobController interface {
	Create(ginContext *gin.Context)
	Update(ginContext *gin.Context)
	Delete(ginContext *gin.Context)
	RequestCompleted(ginContext *gin.Context)
}
