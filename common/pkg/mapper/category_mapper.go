package mapper

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/category/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/category"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func MapCategoryDtoIntoCategoryModel[T *dto.CreateCategoryDto | *dto.UpdateCategoryDto](categoryModel *model.Category, categoryDto T) {
	err := mapstructure.Decode(categoryDto, &categoryModel)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
}
func MapCategoryModelIntoCategoryResponse(categoryModel *model.Category) *category.QueryCategoryResponse {
	var err error
	var queryCategoryResponse category.QueryCategoryResponse
	err = mapstructure.Decode(categoryModel, &queryCategoryResponse)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return &queryCategoryResponse
}

func MapCategoryModelIntoCategoryResponses(categoriesModel []model.Category) category.QueryCategoryResponses {
	var err error

	var queryCategoryResponses = make([]*category.QueryCategoryResponse, 0)
	for _, categoryModel := range categoriesModel {
		var queryCategoryResponse category.QueryCategoryResponse
		err = mapstructure.Decode(categoryModel, &queryCategoryResponse)
		queryCategoryResponses = append(queryCategoryResponses, &queryCategoryResponse)
	}
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	return category.QueryCategoryResponses{
		QueryCategoryResponse: queryCategoryResponses,
	}
}
