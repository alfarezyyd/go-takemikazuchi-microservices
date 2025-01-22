package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-takemikazuchi-api/internal/category"
	"go-takemikazuchi-api/internal/job"
	jobApplication "go-takemikazuchi-api/internal/job_application"
	"go-takemikazuchi-api/internal/middleware"
	"go-takemikazuchi-api/internal/transaction"
	"go-takemikazuchi-api/internal/worker"
)

type ProtectedRoutes struct {
	routerGroup              *gin.RouterGroup
	categoryController       category.Controller
	jobController            job.Controller
	workerController         worker.Controller
	transactionController    transaction.Controller
	jobApplicationController jobApplication.Controller
	viperConfig              *viper.Viper
}

func NewProtectedRoutes(routerGroup *gin.RouterGroup,
	categoryController category.Controller,
	jobController job.Controller,
	viperConfig *viper.Viper,
	workerController worker.Controller,
	jobApplicationController jobApplication.Controller,
) *ProtectedRoutes {
	routerGroup.Use(middleware.AuthMiddleware(viperConfig))
	return &ProtectedRoutes{
		routerGroup:              routerGroup,
		categoryController:       categoryController,
		jobController:            jobController,
		workerController:         workerController,
		viperConfig:              viperConfig,
		jobApplicationController: jobApplicationController,
	}
}

func (protectedRoutes *ProtectedRoutes) Setup() {
	categoryRouterGroup := protectedRoutes.routerGroup.Group("categories")
	categoryRouterGroup.POST("", protectedRoutes.categoryController.Create)
	categoryRouterGroup.PUT("/:id", protectedRoutes.categoryController.Update)
	categoryRouterGroup.DELETE("/:id", protectedRoutes.categoryController.Delete)

	jobRouterGroup := protectedRoutes.routerGroup.Group("jobs")
	jobRouterGroup.POST("", protectedRoutes.jobController.Create)

	workerRouterGroup := protectedRoutes.routerGroup.Group("workers")
	workerRouterGroup.POST("", protectedRoutes.workerController.Register)

	jobApplicationRouterGroup := protectedRoutes.routerGroup.Group("job-applications")
	jobApplicationRouterGroup.POST("", protectedRoutes.jobApplicationController.Apply)
	jobApplicationRouterGroup.POST("/", protectedRoutes.jobApplicationController.SelectApplication)
	jobApplicationRouterGroup.GET("/:jobId", protectedRoutes.jobApplicationController.FindAllApplication)
	jobApplicationRouterGroup.POST("/select", protectedRoutes.jobApplicationController.SelectApplication)
}
