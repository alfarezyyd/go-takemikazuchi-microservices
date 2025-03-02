package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-takemikazuchi-microservices/internal/category"
	"go-takemikazuchi-microservices/internal/job"
	jobApplication "go-takemikazuchi-microservices/internal/job_application"
	"go-takemikazuchi-microservices/internal/middleware"
	reviewFeature "go-takemikazuchi-microservices/internal/review"
	"go-takemikazuchi-microservices/internal/transaction"
	withdrawalFeature "go-takemikazuchi-microservices/internal/withdrawal"
	"go-takemikazuchi-microservices/internal/worker"
)

type ProtectedRoutes struct {
	routerGroup              *gin.RouterGroup
	categoryController       category.Controller
	jobController            job.Controller
	workerController         worker.Controller
	transactionController    transaction.Controller
	jobApplicationController jobApplication.Controller
	reviewController         reviewFeature.Controller
	withdrawalController     withdrawalFeature.Controller
	viperConfig              *viper.Viper
}

func NewProtectedRoutes(routerGroup *gin.RouterGroup,
	categoryController category.Controller,
	jobController job.Controller,
	viperConfig *viper.Viper,
	workerController worker.Controller,
	transactionController transaction.Controller,
	jobApplicationController jobApplication.Controller,
	reviewController reviewFeature.Controller,
	withdrawalController withdrawalFeature.Controller,
) *ProtectedRoutes {
	routerGroup.Use(middleware.AuthMiddleware(viperConfig))
	return &ProtectedRoutes{
		routerGroup:              routerGroup,
		categoryController:       categoryController,
		jobController:            jobController,
		workerController:         workerController,
		viperConfig:              viperConfig,
		transactionController:    transactionController,
		reviewController:         reviewController,
		withdrawalController:     withdrawalController,
		jobApplicationController: jobApplicationController,
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

	workerRouterGroup := protectedRoutes.routerGroup.Group("workers")
	workerRouterGroup.POST("", protectedRoutes.workerController.Register)

	jobApplicationRouterGroup := protectedRoutes.routerGroup.Group("job-applications")
	jobApplicationRouterGroup.POST("", protectedRoutes.jobApplicationController.Apply)
	jobApplicationRouterGroup.POST("/", protectedRoutes.jobApplicationController.SelectApplication)
	jobApplicationRouterGroup.GET("/:jobId", protectedRoutes.jobApplicationController.FindAllApplication)
	jobApplicationRouterGroup.POST("/select", protectedRoutes.jobApplicationController.SelectApplication)

	transactionRouterGroup := protectedRoutes.routerGroup.Group("transactions")
	transactionRouterGroup.POST("", protectedRoutes.transactionController.Create)

	reviewRouterGroup := protectedRoutes.routerGroup.Group("reviews")
	reviewRouterGroup.POST("", protectedRoutes.reviewController.Create)

	withdrawalRouterGroup := protectedRoutes.routerGroup.Group("withdrawals")
	withdrawalRouterGroup.POST("", protectedRoutes.withdrawalController.Create)
	withdrawalRouterGroup.GET("", protectedRoutes.withdrawalController.FindAll)
	withdrawalRouterGroup.PUT("/:withdrawalId", protectedRoutes.withdrawalController.Update)
}
