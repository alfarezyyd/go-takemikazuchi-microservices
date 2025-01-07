package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-takemikazuchi-api/category"
	"go-takemikazuchi-api/job"
	"go-takemikazuchi-api/middleware"
)

type ProtectedRoutes struct {
	routerGroup        *gin.RouterGroup
	categoryController category.Controller
	jobController      job.Controller
	viperConfig        *viper.Viper
}

func NewProtectedRoutes(routerGroup *gin.RouterGroup,
	categoryController category.Controller,
	jobController job.Controller,
	viperConfig *viper.Viper,
) *ProtectedRoutes {
	routerGroup.Use(middleware.AuthMiddleware(viperConfig))
	return &ProtectedRoutes{
		routerGroup:        routerGroup,
		categoryController: categoryController,
		jobController:      jobController,
		viperConfig:        viperConfig,
	}
}

func (routerGroup *ProtectedRoutes) Setup() {
	categoryRouterGroup := routerGroup.routerGroup.Group("categories")
	categoryRouterGroup.POST("", routerGroup.categoryController.Create)
	categoryRouterGroup.PUT("/:id", routerGroup.categoryController.Update)
	categoryRouterGroup.DELETE("/:id", routerGroup.categoryController.Delete)

	jobRouterGroup := routerGroup.routerGroup.Group("jobs")
	jobRouterGroup.POST("", routerGroup.jobController.Create)
}
