package mapper

import (
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	"github.com/go-viper/mapstructure/v2"
	"net/http"
)

func MapWithdrawalDtoIntoWithdrawalModel(withdrawalDto *dto.CreateWithdrawalDto) *model.Withdrawal {
	var withdrawal *model.Withdrawal
	err := mapstructure.Decode(withdrawalDto, &withdrawal)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	return withdrawal
}
