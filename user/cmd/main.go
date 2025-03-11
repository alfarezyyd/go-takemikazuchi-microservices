package main

import (
	"context"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/middleware"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/handler"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/service"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	serviceName = "userService"
	httpAddr    = ":7001"
	grpcAddr    = ":10001"
	consulAddr  = ":8500"
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
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(middleware.RecoveryInterceptor),
	)

	validatorInstance, engTranslator := configs.InitializeValidator()
	mailerService := configs.NewMailerService(viperConfig)
	identityProvider := configs.NewIdentityProvider(viperConfig)
	validatorService := validatorFeature.NewService(validatorInstance, engTranslator)

	userRepository := repository.NewUserRepository()
	addressRepository := repository.NewUserAddressRepository()
	nominatimBaseUrl := viperConfig.GetString("NOMINATIM_BASE_URL")
	nominatimHttpClient := configs.NewHttpClient(nil, &nominatimBaseUrl)
	userService := service.NewUserService(validatorService, userRepository, databaseConnection, mailerService, identityProvider, viperConfig)
	userAddressService := service.NewUserAddressServiceImpl(userRepository, databaseConnection, addressRepository, validatorService, nominatimHttpClient)
	handler.NewUserHandler(grpcServer, userService)
	handler.NewUserAddressHandler(grpcServer, userAddressService)
	fmt.Println("Serving gRPC server at " + grpcAddr)
	err = grpcServer.Serve(tcpListener)

	if err != nil {
		log.Fatalf("Failed to serve gRPC connection: %v", err)
	}
}
