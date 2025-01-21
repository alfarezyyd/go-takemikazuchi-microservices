package worker

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-api/internal/worker/dto"
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
	err := ginContext.ShouldBindBodyWithJSON(createWorkerDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	identityCard, _ := ginContext.FormFile("identityCard")
	policeCertificate, _ := ginContext.FormFile("policeCertificate")
	driverLicense, _ := ginContext.FormFile("driverLicense")

	createWorkerWalletDto := &dto.CreateWorkerWalletDocumentDto{
		IdentityCard:      identityCard,
		PoliceCertificate: policeCertificate,
		DriverLicense:     driverLicense,
	}
	workerHandler.WorkerService.Create(&createWorkerDto, createWorkerWalletDto)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Success", nil))
}
