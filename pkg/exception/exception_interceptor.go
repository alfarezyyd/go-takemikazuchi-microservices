package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-api/web"
	"net/http"
)

func Interceptor() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		defer func() {
			if occurredError := recover(); occurredError != nil {
				fmt.Println("panic occurred", occurredError)
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
