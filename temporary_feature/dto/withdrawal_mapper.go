package dto

import (
	"errors"
	"github.com/go-viper/mapstructure/v2"
	"go-takemikazuchi-microservices/internal/model"
	"go-takemikazuchi-microservices/internal/withdrawal/dto"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"net/http"
)

func MapCreateWithdrawalDtoIntoWithdrawalModel(createWithdrawalDto *dto.CreateWithdrawalDto, withdrawModel *model.Withdrawal) {
	err := mapstructure.Decode(createWithdrawalDto, withdrawModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
}
