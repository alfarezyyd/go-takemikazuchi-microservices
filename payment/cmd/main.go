package main

import (
	"context"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/middleware"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/internal/handler"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/internal/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/internal/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	serviceName = "paymentService"
	httpAddr    = ":7006"
	consulAddr  = ":8500"
	grpcAddr    = ":10006"
)

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
	tcpListener, err := net.Listen("tcp", grpcAddr)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middleware.RecoveryInterceptor),
	)

	if err != nil {
		log.Fatalf("Failed to create gRPC connection: %v", err)
	}
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
	validatorInstance, engTranslator := configs.InitializeValidator()

	midtransService := configs.NewMidtransService(viperConfig)
	midtransClient := midtransService.InitializeMidtransConfiguration()
	transactionRepository := repository.NewTransactionRepository()
	rabbitMQURL := "amqp://guest:guest@localhost:5672/"
	queueName := "OrderUpdate"
	rabbitMQ, err := configs.NewRabbitMQ(rabbitMQURL, queueName)
	if err != nil {
		log.Fatalf("Failed to create rabbitMQ instance: %v", err)
	}
	transactionService := service.NewTransactionService(validatorInstance, engTranslator, databaseConnection, midtransClient, transactionRepository, viperConfig, consulServiceRegistry, rabbitMQ)
	handler.NewTransactionHandler(grpcServer, transactionService)
	fmt.Println("gRPC server listening on " + tcpListener.Addr().String())
	err = grpcServer.Serve(tcpListener)
	if err != nil {
		log.Fatalf("Failed to serve gRPC connection: %v", err)
	}
}
