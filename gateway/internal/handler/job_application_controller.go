package handler

import "github.com/gin-gonic/gin"

type JobApplicationController interface {
	FindAllApplication(ginContext *gin.Context)
	SelectApplication(ginContext *gin.Context)
	Apply(ginContext *gin.Context)
}
