package handler

import (
	"context"
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/worker"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type WorkerHandler struct {
	serviceRegistry discovery.ServiceRegistry
}

func NewWorkerHandler(serviceRegistry discovery.ServiceRegistry) *WorkerHandler {
	return &WorkerHandler{
		serviceRegistry: serviceRegistry,
	}
}

func (workerHandler *WorkerHandler) Register(ginContext *gin.Context) {
	var createWorkerDto dto.CreateWorkerDto
	err := ginContext.ShouldBind(&createWorkerDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	//identityCard, _ := ginContext.FormFile("identityCard")
	//policeCertificate, _ := ginContext.FormFile("policeCertificate")
	//driverLicense, _ := ginContext.FormFile("driverLicense")
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	//createWorkerWalletDto := &workerResourceDto.CreateWorkerWalletDocumentDto{
	//	IdentityCard:      identityCard,
	//	PoliceCertificate: policeCertificate,
	//	DriverLicense:     driverLicense,
	//}
	timeoutBackground, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	gRCPClientConnection, err := discovery.ServiceConnection(timeoutBackground, "workerService", workerHandler.serviceRegistry)
	workerServiceClient := worker.NewWorkerServiceClient(gRCPClientConnection)
	createWorkerRequest := mapper.MapCreateWorkerDtoIntoCreateWorkerRequest(createWorkerDto)
	createWorkerRequest.UserJwtClaim = mapper.MapUserJwtClaimIntoUserJwtClaimGrpc(userJwtClaim)
	_, err = workerServiceClient.Create(timeoutBackground, createWorkerRequest)
	exception.ParseGrpcError(err)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Register for worker success", nil))
}
