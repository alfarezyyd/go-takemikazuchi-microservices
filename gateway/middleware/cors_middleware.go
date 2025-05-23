package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func CorsMiddleware() gin.HandlerFunc {
	// Define allowed origins as a comma-separated string
	originsString := "http://localhost:3001,http://localhost:3000"
	var allowedOrigins []string
	allowedOrigins = strings.Split(originsString, ",")

	// Return the actual middleware handler function
	return func(c *gin.Context) {
		// Function to check if a given origin is allowed
		isOriginAllowed := func(origin string, allowedOrigins []string) bool {
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					return true
				}
			}
			return false
		}

		// Get the Origin header from the request
		origin := c.Request.Header.Get("Origin")
		fmt.Println("CORS TRIGGERED")
		// Check if the origin is allowed
		if isOriginAllowed(origin, allowedOrigins) {
			fmt.Println("CorsMiddleware")
			// If the origin is allowed, set CORS headers in the response
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		}

		// Handle preflight OPTIONS requests by aborting with status 204
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Call the next handler
		c.Next()
	}
}
