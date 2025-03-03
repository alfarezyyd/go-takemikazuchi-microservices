package mapper

import (
	"errors"
	"github.com/go-viper/mapstructure/v2"
	"go-takemikazuchi-microservices/internal/model"
	"go-takemikazuchi-microservices/internal/worker/dto"
	workerWalletDto "go-takemikazuchi-microservices/internal/worker_wallet/dto"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"net/http"
)

func MapCreateWorkerDtoIntoWorkerModel(workerModel *model.Worker, createWorkerDto *dto.CreateWorkerDto) {
	err := mapstructure.Decode(createWorkerDto, workerModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
}

func MapCreateWorkerWalletDtoIntoWorkerWalletModel(workerWalletModel *model.WorkerWallet, createWorkerWalletDto *workerWalletDto.CreateWorkerWalletDto) {
	err := mapstructure.Decode(createWorkerWalletDto, workerWalletModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
}
