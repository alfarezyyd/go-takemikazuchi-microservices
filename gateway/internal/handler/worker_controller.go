package handler

import "github.com/gin-gonic/gin"

type WorkerController interface {
	Register(ginContext *gin.Context)
}
