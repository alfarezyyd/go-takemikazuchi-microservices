package job

import "github.com/gin-gonic/gin"

type Controller interface {
	Create(ginContext *gin.Context)
	Update(ginContext *gin.Context)
	Delete(ginContext *gin.Context)
	RequestCompleted(ginContext *gin.Context)
}
