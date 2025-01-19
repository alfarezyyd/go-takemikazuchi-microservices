package mapper

import (
	"github.com/go-viper/mapstructure/v2"
	dto2 "go-takemikazuchi-api/internal/category/dto"
	"go-takemikazuchi-api/internal/model"
	exception2 "go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"net/http"
)

func MapCategoryDtoIntoCategoryModel[T *dto2.CreateCategoryDto | *dto2.UpdateCategoryDto](categoryModel *model.Category, categoryDto T) {
	err := mapstructure.Decode(categoryDto, &categoryModel)
	helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
}
