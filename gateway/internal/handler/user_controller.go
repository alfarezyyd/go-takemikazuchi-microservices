package handler

import "github.com/gin-gonic/gin"

type UserController interface {
	Login(ginContext *gin.Context)
	Register(ginContext *gin.Context)
	GenerateOneTimePassword(ginContext *gin.Context)
	VerifyOneTimePassword(ginContext *gin.Context)
}
