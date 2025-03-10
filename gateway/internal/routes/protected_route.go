package routes

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/middleware"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/gateway/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ProtectedRoutes struct {
	routerGroup              *gin.RouterGroup
	categoryController       handler.CategoryController
	jobController            handler.JobController
	viperConfig              *viper.Viper
	workerController         handler.WorkerController
	jobApplicationController handler.JobApplicationController
}

func NewProtectedRoutes(routerGroup *gin.RouterGroup,
	categoryController handler.CategoryController,
	jobController handler.JobController,
	workerController handler.WorkerController,
	jobApplicationController handler.JobApplicationController,
	viperConfig *viper.Viper,
) *ProtectedRoutes {
	routerGroup.Use(middleware.AuthMiddleware(viperConfig))
	return &ProtectedRoutes{
		routerGroup:              routerGroup,
		jobController:            jobController,
		categoryController:       categoryController,
		workerController:         workerController,
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

}
