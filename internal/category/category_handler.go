package category

import (
	"github.com/gin-gonic/gin"
	dto2 "go-takemikazuchi-api/internal/category/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
	exception2 "go-takemikazuchi-api/pkg/exception"
	helper2 "go-takemikazuchi-api/pkg/helper"
	"net/http"
)

type Handler struct {
	categoryService Service
}

func NewHandler(categoryService Service) *Handler {
	return &Handler{
		categoryService: categoryService,
	}
}

func (categoryHandler *Handler) Create(ginContext *gin.Context) {
	var categoryCreateDto dto2.CreateCategoryDto
	err := ginContext.ShouldBindBodyWithJSON(&categoryCreateDto)
	helper2.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	clientError := categoryHandler.categoryService.HandleCreate(userJwtClaim, &categoryCreateDto)
	helper2.CheckErrorOperation(clientError.GetRawError(), clientError)
	ginContext.JSON(http.StatusCreated, helper2.WriteSuccess("Category has been created", nil))
}

func (categoryHandler *Handler) Update(ginContext *gin.Context) {
	var updateCategoryDto dto2.UpdateCategoryDto
	err := ginContext.ShouldBindBodyWithJSON(&updateCategoryDto)
	helper2.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	categoryId := ginContext.Param("id")
	clientError := categoryHandler.categoryService.HandleUpdate(categoryId, userJwtClaim, &updateCategoryDto)
	helper2.CheckErrorOperation(clientError.GetRawError(), clientError)
	ginContext.JSON(http.StatusOK, helper2.WriteSuccess("Category has been updated", nil))
}

func (categoryHandler *Handler) Delete(ginContext *gin.Context) {
	categoryId := ginContext.Param("id")
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	clientError := categoryHandler.categoryService.HandleDelete(categoryId, userJwtClaim)
	helper2.CheckErrorOperation(clientError.GetRawError(), clientError)
	ginContext.JSON(http.StatusOK, helper2.WriteSuccess("Category has been deleted", nil))
}
