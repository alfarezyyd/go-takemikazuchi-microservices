package routes

import (
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-microservices/internal/user"
)

type AuthenticationRoutes struct {
	routerGroup    *gin.RouterGroup
	userController user.Controller
}

func NewAuthenticationRoutes(routerGroup *gin.RouterGroup, userController user.Controller) *AuthenticationRoutes {
	return &AuthenticationRoutes{
		routerGroup:    routerGroup.Group("authentication"),
		userController: userController,
	}
}

func (routerGroup *AuthenticationRoutes) Setup() {
	routerGroup.routerGroup.GET("google", routerGroup.userController.LoginWithGoogle)
	routerGroup.routerGroup.GET("google/callback", routerGroup.userController.GoogleProviderCallback)
	routerGroup.routerGroup.POST("login", routerGroup.userController.Login)
	routerGroup.routerGroup.POST("register", routerGroup.userController.Register)
	routerGroup.routerGroup.POST("generate-otp", routerGroup.userController.GenerateOneTimePassword)
	routerGroup.routerGroup.POST("verify-otp", routerGroup.userController.VerifyOneTimePassword)
}
