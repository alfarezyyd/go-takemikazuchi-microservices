package exception

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Interceptor() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		defer func() {
			if occurredError := recover(); occurredError != nil {
				// Check if it's our custom error
				if clientError, ok := occurredError.(*ClientError); ok {
					ginContext.AbortWithStatusJSON(
						clientError.StatusCode,
						web.NewResponseContract(false, clientError.Message, nil, &clientError.Trace),
					)
					return
				}

				// Unknown error
				ginContext.AbortWithStatusJSON(
					http.StatusInternalServerError,
					web.NewResponseContract(false, "Internal server error", nil, nil),
				)
			}
		}()
		ginContext.Next()
	}
}
