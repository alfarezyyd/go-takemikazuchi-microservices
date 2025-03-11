package main

import (
	"context"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/middleware"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/handler"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	serviceName = "jobService"
	httpAddr    = ":7003"
	grpcAddr    = ":10003"
	consulAddr  = ":8500"
	rabbitMQURL = "amqp://guest:guest@localhost:5672/"
	queueName   = "order_update"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	consulServiceRegistry, err := discovery.NewRegistry(consulAddr)
	if err != nil {
		panic(err)
	}

	serviceId := discovery.GenerateInstanceID(serviceName)
	ctx := context.Background()
	if err := consulServiceRegistry.Register(ctx, serviceId, serviceName, grpcAddr); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := consulServiceRegistry.HealthCheck(serviceId, serviceName); err != nil {
				log.Fatal("failed to health check")
			}
			time.Sleep(time.Second * 1)
		}
	}()
	defer consulServiceRegistry.Deregister(ctx, serviceId, serviceName)

	mux := http.NewServeMux()

	go func() {
		log.Printf("Starting HTTP server at %s", httpAddr)
		if err := http.ListenAndServe(httpAddr, mux); err != nil {
			log.Fatal("failed to start http server")
		}
	}()

	viperConfig := viper.New()
	viperConfig.SetConfigFile(".env")
	viperConfig.AddConfigPath(".")
	viperConfig.AutomaticEnv()
	viperConfig.ReadInConfig()

	// Database Initialization
	databaseCredentials := &configs.DatabaseCredentials{
		DatabaseHost:     viperConfig.GetString("DATABASE_HOST"),
		DatabasePort:     viperConfig.GetString("DATABASE_PORT"),
		DatabaseName:     viperConfig.GetString("DATABASE_NAME"),
		DatabasePassword: viperConfig.GetString("DATABASE_PASSWORD"),
		DatabaseUsername: viperConfig.GetString("DATABASE_USERNAME"),
	}

	databaseInstance := configs.NewDatabaseConnection(databaseCredentials)
	databaseConnection := databaseInstance.GetDatabaseConnection()

	tcpListener, err := net.Listen("tcp", grpcAddr)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middleware.RecoveryInterceptor),
	)

	jobRepository := repository.NewJobRepository()
	jobApplicationRepository := repository.NewJobApplicationRepository()
	validatorInstance, engTranslator := configs.InitializeValidator()
	validatorService := validatorFeature.NewService(validatorInstance, engTranslator)

	// 1. Buat RabbitMQ Consumer
	consumer, err := configs.NewConsumer(rabbitMQURL, queueName)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
	}
	jobService := service.NewJobService(jobRepository, databaseConnection, nil, validatorService, consulServiceRegistry, consumer)
	jobApplicationService := service.NewJobApplicationService(validatorInstance, engTranslator, jobApplicationRepository, databaseConnection, jobRepository, validatorService, consulServiceRegistry)
	handler.NewJobHandler(grpcServer, jobService)
	handler.NewJobApplicationHandler(grpcServer, jobApplicationService)

	fmt.Println("Serving gRPC server at " + grpcAddr)
	err = grpcServer.Serve(tcpListener)

	if err != nil {
		log.Fatalf("Failed to serve gRPC connection: %v", err)
	}
}
