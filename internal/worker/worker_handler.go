package worker

import (
	"errors"
	"github.com/gin-gonic/gin"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/internal/worker/dto"
	workerResourceDto "go-takemikazuchi-api/internal/worker_resource/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"net/http"
)

type Handler struct {
	WorkerService Service
}

func NewHandler(workerService Service) *Handler {
	return &Handler{
		WorkerService: workerService,
	}
}

func (workerHandler *Handler) Register(ginContext *gin.Context) {
	var createWorkerDto dto.CreateWorkerDto
	err := ginContext.ShouldBind(&createWorkerDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	identityCard, _ := ginContext.FormFile("identityCard")
	policeCertificate, _ := ginContext.FormFile("policeCertificate")
	driverLicense, _ := ginContext.FormFile("driverLicense")
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	createWorkerWalletDto := &workerResourceDto.CreateWorkerWalletDocumentDto{
		IdentityCard:      identityCard,
		PoliceCertificate: policeCertificate,
		DriverLicense:     driverLicense,
	}
	workerHandler.WorkerService.Create(userJwtClaim, &createWorkerDto, createWorkerWalletDto)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Success", nil))
}
