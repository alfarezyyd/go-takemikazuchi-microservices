package mapper

import (
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	workerWallet "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/worker_wallet"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/pkg/dto"
	"github.com/go-viper/mapstructure/v2"
	"net/http"
)

func MapWorkerWalletModelIntoWorkerWalletResponse(workerModel *model.WorkerWallet) *dto.ResponseWorkerWalletDto {
	var responseWorkerWalletDto dto.ResponseWorkerWalletDto
	err := mapstructure.Decode(workerModel, &responseWorkerWalletDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &responseWorkerWalletDto
}

func MapWorkerWalletResponseIntoWorkerWalletGrpcResponse(workerWalletResponseDto *dto.ResponseWorkerWalletDto) *workerWallet.WorkerWalletResponse {
	var workerWalletResponse workerWallet.WorkerWalletResponse
	err := mapstructure.Decode(workerWalletResponseDto, &workerWalletResponse)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &workerWalletResponse
}

func MapCreateWorkerWalletDtoIntoWorkerWalletModel(workerWalletModel *model.WorkerWallet, createWorkerWalletDto *dto.CreateWorkerWalletDto) {
	err := mapstructure.Decode(createWorkerWalletDto, workerWalletModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
}
