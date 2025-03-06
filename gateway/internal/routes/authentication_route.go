package routes

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/gateway/internal/handler"
	"github.com/gin-gonic/gin"
)

type AuthenticationRoutes struct {
	routerGroup    *gin.RouterGroup
	userController handler.UserController
}

func NewAuthenticationRoutes(routerGroup *gin.RouterGroup, userController handler.UserController) *AuthenticationRoutes {
	return &AuthenticationRoutes{
		routerGroup:    routerGroup.Group("authentication"),
		userController: userController,
	}
}

func (routerGroup *AuthenticationRoutes) Setup() {
	routerGroup.routerGroup.POST("login", routerGroup.userController.Login)
	routerGroup.routerGroup.POST("register", routerGroup.userController.Register)
	routerGroup.routerGroup.POST("generate-otp", routerGroup.userController.GenerateOneTimePassword)
	routerGroup.routerGroup.POST("verify-otp", routerGroup.userController.VerifyOneTimePassword)
}
