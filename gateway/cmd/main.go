package main

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/gateway/internal/handler"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/gateway/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	serviceName = "gatewayService"
	httpAddr    = ":7000"
	consulAddr  = ":8500"
	grpcAddr    = ":10001"
)

func main() {
	consulServiceRegistry, err := discovery.NewRegistry(consulAddr)
	if err != nil {
		panic(err)
	}

	serviceId := discovery.GenerateInstanceID(serviceName)
	ctx := context.Background()
	if err := consulServiceRegistry.Register(ctx, serviceId, serviceName, httpAddr); err != nil {
		panic(err)
	}

	defer consulServiceRegistry.Deregister(ctx, serviceId, serviceName)
	traceProvider, err := configs.StartTracing()
	if err != nil {
		log.Fatalf("traceprovider: %v", err)
	}
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			log.Fatalf("traceprovider: %v", err)
		}
	}()

	_ = traceProvider.Tracer("suruhAkuAja")
	if err != nil {
		log.Fatalf("Failed to create gRPC connection: %v", err)
	}
	viperConfig := viper.New()
	viperConfig.SetConfigFile(".env")
	viperConfig.AddConfigPath(".")
	viperConfig.AutomaticEnv()
	viperConfig.ReadInConfig()

	ginEngine := gin.Default()
	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // atau domain frontend kamu
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,

		MaxAge: 12 * time.Hour,
	}))
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(exception.Interceptor())

	rootRouterGroup := ginEngine.Group("/")
	userHandler := handler.NewUserHandler(consulServiceRegistry)
	categoryHandler := handler.NewCategoryHandler(consulServiceRegistry)
	jobHandler := handler.NewJobHandler(consulServiceRegistry)
	workerHandler := handler.NewWorkerHandler(consulServiceRegistry)
	authenticationRoutes := routes.NewAuthenticationRoutes(rootRouterGroup, userHandler)
	jobApplicationHandler := handler.NewJobApplicationHandler(consulServiceRegistry)
	transactionHandler := handler.NewTransactionHandler(consulServiceRegistry)
	withdrawalHandler := handler.NewWithdrawalHandler(consulServiceRegistry)
	protectedRoutes := routes.NewProtectedRoutes(rootRouterGroup, categoryHandler, jobHandler, workerHandler, jobApplicationHandler, viperConfig, transactionHandler, withdrawalHandler)
	publicRoutes := routes.NewPublicRoutes(ginEngine.Group("/"), transactionHandler)
	authenticationRoutes.Setup()
	protectedRoutes.Setup()
	publicRoutes.Setup()
	ginError := ginEngine.Run(":8080")
	if ginError != nil {
		panic(ginError)
	}
}
