package main

import (
	"fmt"
	categoryHandler "github.com/alfarezyyd/go-takemikazuchi-microservices/category/internal/category/handler"
	categoryRepository "github.com/alfarezyyd/go-takemikazuchi-microservices/category/internal/category/repository"
	categoryService "github.com/alfarezyyd/go-takemikazuchi-microservices/category/internal/category/service"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/middleware"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func main() {
	grpcConnection, err := grpc.NewClient("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	tcpListener, err := net.Listen("tcp", ":9000")
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middleware.RecoveryInterceptor),
	)

	if err != nil {
		log.Fatalf("Failed to create gRPC connection: %v", err)
	}
	defer grpcConnection.Close()
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
	fmt.Println(databaseCredentials)
	databaseInstance := configs.NewDatabaseConnection(databaseCredentials)
	databaseConnection := databaseInstance.GetDatabaseConnection()
	categoryServiceImpl := categoryRepository.NewRepository()
	validatorInstance, engTranslator := configs.InitializeValidator()

	validatorService := validatorFeature.NewService(validatorInstance, engTranslator)
	categoryServiceInstance := categoryService.NewService(categoryServiceImpl, databaseConnection, validatorService)
	categoryHandler.NewCategoryHandler(grpcServer, categoryServiceInstance)
	fmt.Println("gRPC server listening on " + tcpListener.Addr().String())
	err = grpcServer.Serve(tcpListener)
	if err != nil {
		log.Fatalf("Failed to serve gRPC connection: %v", err)
	}
}
