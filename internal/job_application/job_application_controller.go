package job_application

import "github.com/gin-gonic/gin"

type Controller interface {
	Apply(ginContext *gin.Context)
}
