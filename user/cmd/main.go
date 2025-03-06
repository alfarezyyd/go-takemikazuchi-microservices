package main

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/user/handler"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/user/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/user/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"time"
)

var (
	serviceName = "userService"
	httpAddr    = ":8080"
	consulAddr  = ":8500"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	consulServiceRegistry, err := discovery.NewRegistry(consulAddr, serviceName)
	if err != nil {
		panic(err)
	}

	serviceId := discovery.GenerateInstanceID(serviceName)
	ctx := context.Background()
	if err := consulServiceRegistry.Register(ctx, serviceId, serviceName, httpAddr); err != nil {
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

	grpcConnection, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer grpcConnection.Close()

	if err != nil {
		log.Fatalf("Failed to create gRPC connection: %v", err)
	}
	tcpListener, err := net.Listen("tcp", ":9000")
	grpcServer := grpc.NewServer()
	err = grpcServer.Serve(tcpListener)
	userRepository := repository.NewRepository()
	validatorInstance, engTranslator := configs.InitializeValidator()
	mailerService := configs.NewMailerService(viperConfig)
	identityProvider := configs.NewIdentityProvider(viperConfig)
	validatorService := validatorFeature.NewService(validatorInstance, engTranslator)

	newService := service.NewService(validatorService, userRepository, databaseConnection, mailerService, identityProvider, viperConfig)
	handler.NewUserHandler(grpcServer, newService)
	err = grpcServer.Serve(tcpListener)
	if err != nil {
		log.Fatalf("Failed to serve gRPC connection: %v", err)
	}
}
