package handler

import "github.com/gin-gonic/gin"

type WithdrawalController interface {
	//FindAll(ginContext *gin.Context)
	Create(ginContext *gin.Context)
	//Update(ginContext *gin.Context)
}
