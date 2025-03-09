package mapper

import (
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/worker"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/pkg/dto"
	"github.com/go-viper/mapstructure/v2"
	"net/http"
)

func MapCreateWorkerDtoIntoWorkerModel(workerModel *model.Worker, createWorkerDto *dto.CreateWorkerDto) {
	err := mapstructure.Decode(createWorkerDto, workerModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
}

func MapCreateWorkerWalletDtoIntoWorkerWalletModel(workerWalletModel *model.WorkerWallet, createWorkerWalletDto *dto.CreateWorkerWalletDto) {
	err := mapstructure.Decode(createWorkerWalletDto, workerWalletModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
}

func MapCreateWorkerDtoIntoCreateWorkerRequest(createWorkerDto dto.CreateWorkerDto) *worker.CreateWorkerRequest {
	var createWorkerRequest worker.CreateWorkerRequest
	err := mapstructure.Decode(createWorkerDto, &createWorkerRequest)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	return &createWorkerRequest
}

func MapCreateWorkerRequestIntoCreateWorkerDto(createWorkerRequest *worker.CreateWorkerRequest) *dto.CreateWorkerDto {
	var createWorkerRequestDto dto.CreateWorkerDto
	err := mapstructure.Decode(createWorkerRequest, &createWorkerRequestDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	return &createWorkerRequestDto
}
