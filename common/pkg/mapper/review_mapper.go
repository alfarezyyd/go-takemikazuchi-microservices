package mapper

import (
	"errors"
	"github.com/go-viper/mapstructure/v2"
	"go-takemikazuchi-microservices/internal/model"
	"go-takemikazuchi-microservices/internal/review/dto"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"net/http"
)

func MapCreateReviewDtoIntoReviewModel(createReviewDto *dto.CreateReviewDto, reviewModel *model.Review) {
	err := mapstructure.Decode(createReviewDto, reviewModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("error when parsing")))
}
