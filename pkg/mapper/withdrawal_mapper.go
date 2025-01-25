package mapper

import (
	"errors"
	"github.com/go-viper/mapstructure/v2"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/internal/withdrawal/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"net/http"
)

func MapCreateWithdrawalDtoIntoWithdrawalModel(createWithdrawalDto *dto.CreateWithdrawalDto, withdrawModel *model.Withdrawal) {
	err := mapstructure.Decode(createWithdrawalDto, withdrawModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
}
