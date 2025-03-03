package handler

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices-category/internal/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-category/internal/service"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	categoryService service.Service
}

func NewHandler(categoryService service.Service) *Handler {
	return &Handler{
		categoryService: categoryService,
	}
}

func (categoryHandler *Handler) FindAll(ginContext *gin.Context) {
	categoriesResponseDto := categoryHandler.categoryService.FindAll()
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Success", categoriesResponseDto))
}

func (categoryHandler *Handler) Create(ginContext *gin.Context) {
	var categoryCreateDto dto.CreateCategoryDto
	err := ginContext.ShouldBindBodyWithJSON(&categoryCreateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	clientError := categoryHandler.categoryService.HandleCreate(userJwtClaim, &categoryCreateDto)
	helper.CheckErrorOperation(clientError.GetRawError(), clientError)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Category has been created", nil))
}

func (categoryHandler *Handler) Update(ginContext *gin.Context) {
	var updateCategoryDto dto.UpdateCategoryDto
	err := ginContext.ShouldBindBodyWithJSON(&updateCategoryDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	categoryId := ginContext.Param("id")
	clientError := categoryHandler.categoryService.HandleUpdate(categoryId, userJwtClaim, &updateCategoryDto)
	helper.CheckErrorOperation(clientError.GetRawError(), clientError)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Category has been updated", nil))
}

func (categoryHandler *Handler) Delete(ginContext *gin.Context) {
	categoryId := ginContext.Param("id")
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	clientError := categoryHandler.categoryService.HandleDelete(categoryId, userJwtClaim)
	helper.CheckErrorOperation(clientError.GetRawError(), clientError)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Category has been deleted", nil))
}
