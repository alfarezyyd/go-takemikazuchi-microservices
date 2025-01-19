package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-takemikazuchi-api/internal/category"
	"go-takemikazuchi-api/internal/job"
	"go-takemikazuchi-api/internal/job_application"
	"go-takemikazuchi-api/internal/middleware"
	"go-takemikazuchi-api/internal/transaction"
)

type ProtectedRoutes struct {
	routerGroup              *gin.RouterGroup
	categoryController       category.Controller
	jobController            job.Controller
	transactionController    transaction.Controller
	jobApplicationController job_application.Controller
	viperConfig              *viper.Viper
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

func (protectedRoutes *ProtectedRoutes) Setup() {
	categoryRouterGroup := protectedRoutes.routerGroup.Group("categories")
	categoryRouterGroup.POST("", protectedRoutes.categoryController.Create)
	categoryRouterGroup.PUT("/:id", protectedRoutes.categoryController.Update)
	categoryRouterGroup.DELETE("/:id", protectedRoutes.categoryController.Delete)

	jobRouterGroup := protectedRoutes.routerGroup.Group("jobs")
	jobRouterGroup.POST("", protectedRoutes.jobController.Create)
}
