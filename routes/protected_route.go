package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-takemikazuchi-api/category"
	"go-takemikazuchi-api/middleware"
)

type ProtectedRoutes struct {
	routerGroup        *gin.RouterGroup
	categoryController category.Controller
	viperConfig        *viper.Viper
}

func NewProtectedRoutes(routerGroup *gin.RouterGroup, categoryController category.Controller, viperConfig *viper.Viper) *ProtectedRoutes {
	routerGroup.Use(middleware.AuthMiddleware(viperConfig))
	return &ProtectedRoutes{
		routerGroup:        routerGroup.Group("categories"),
		categoryController: categoryController,
		viperConfig:        viperConfig,
	}
}

func (routerGroup *ProtectedRoutes) Setup() {
	routerGroup.routerGroup.POST("", routerGroup.categoryController.Create)
	routerGroup.routerGroup.PUT("/:id", routerGroup.categoryController.Update)
	routerGroup.routerGroup.DELETE("/:id", routerGroup.categoryController.Delete)
}
