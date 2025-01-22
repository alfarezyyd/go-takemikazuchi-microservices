package job_application

import "github.com/gin-gonic/gin"

type Controller interface {
	FindAllApplication(ginContext *gin.Context)
	Apply(ginContext *gin.Context)
}
