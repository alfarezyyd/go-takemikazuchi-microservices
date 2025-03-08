package routes

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/middleware"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/gateway/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ProtectedRoutes struct {
	routerGroup        *gin.RouterGroup
	categoryController handler.CategoryController

	viperConfig *viper.Viper
}

func NewProtectedRoutes(routerGroup *gin.RouterGroup,
	categoryController handler.CategoryController,
	viperConfig *viper.Viper,

) *ProtectedRoutes {
	routerGroup.Use(middleware.AuthMiddleware(viperConfig))
	return &ProtectedRoutes{
		routerGroup:        routerGroup,
		categoryController: categoryController,
	}
}

func (protectedRoutes *ProtectedRoutes) Setup() {
	categoryRouterGroup := protectedRoutes.routerGroup.Group("categories")
	categoryRouterGroup.POST("", protectedRoutes.categoryController.Create)
	categoryRouterGroup.GET("", protectedRoutes.categoryController.FindAll)
	categoryRouterGroup.PUT("/:id", protectedRoutes.categoryController.Update)
	categoryRouterGroup.DELETE("/:id", protectedRoutes.categoryController.Delete)

}
