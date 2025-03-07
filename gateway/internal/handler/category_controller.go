package handler

import "github.com/gin-gonic/gin"

type CategoryController interface {
	FindAll(ginContext *gin.Context)
	Create(ginContext *gin.Context)
	Update(ginContext *gin.Context)
	Delete(ginContext *gin.Context)
}
