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
	jobController      handler.JobController
	viperConfig        *viper.Viper
}

func NewProtectedRoutes(routerGroup *gin.RouterGroup,
	categoryController handler.CategoryController,
	jobController handler.JobController,
	viperConfig *viper.Viper,
) *ProtectedRoutes {
	routerGroup.Use(middleware.AuthMiddleware(viperConfig))
	return &ProtectedRoutes{
		routerGroup:        routerGroup,
		jobController:      jobController,
		categoryController: categoryController,
	}
}

func (protectedRoutes *ProtectedRoutes) Setup() {
	categoryRouterGroup := protectedRoutes.routerGroup.Group("categories")
	categoryRouterGroup.POST("", protectedRoutes.categoryController.Create)
	categoryRouterGroup.GET("", protectedRoutes.categoryController.FindAll)
	categoryRouterGroup.PUT("/:id", protectedRoutes.categoryController.Update)
	categoryRouterGroup.DELETE("/:id", protectedRoutes.categoryController.Delete)

	jobRouterGroup := protectedRoutes.routerGroup.Group("jobs")
	jobRouterGroup.POST("", protectedRoutes.jobController.Create)
	jobRouterGroup.PUT("/:jobId", protectedRoutes.jobController.Update)
	jobRouterGroup.POST("/completed/:jobId", protectedRoutes.jobController.RequestCompleted)
}
