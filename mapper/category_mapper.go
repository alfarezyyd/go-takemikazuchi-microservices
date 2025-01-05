package mapper

import (
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"go-takemikazuchi-api/category/dto"
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/helper"
	"go-takemikazuchi-api/model"
	"net/http"
)

func MapCategoryDtoIntoCategoryModel[T *dto.CreateCategoryDto | *dto.UpdateCategoryDto](categoryModel *model.Category, categoryDto T) {
	err := mapstructure.Decode(categoryDto, &categoryModel)
	fmt.Println(err)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest))
}
