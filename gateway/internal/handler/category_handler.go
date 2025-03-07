package handler

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/category/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/category"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"time"
)

type CategoryHandler struct {
	serviceRegistry discovery.ServiceRegistry
}

func NewCategoryHandler(serviceRegistry discovery.ServiceRegistry) *CategoryHandler {
	return &CategoryHandler{
		serviceRegistry: serviceRegistry,
	}
}

func (categoryHandler *CategoryHandler) FindAll(ginContext *gin.Context) {
	timeoutBackground, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	gRCPClientConnection, err := discovery.ServiceConnection(timeoutBackground, "categoryService", categoryHandler.serviceRegistry)
	categoryServiceClient := category.NewCategoryServiceClient(gRCPClientConnection)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	categoriesResponseDto, err := categoryServiceClient.FindAll(timeoutBackground, &emptypb.Empty{})
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Success", categoriesResponseDto))
}

func (categoryHandler *CategoryHandler) Create(ginContext *gin.Context) {
	var categoryCreateDto dto.CreateCategoryDto
	err := ginContext.ShouldBindBodyWithJSON(&categoryCreateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	timeoutBackground, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	gRCPClientConnection, err := discovery.ServiceConnection(timeoutBackground, "categoryService", categoryHandler.serviceRegistry)
	categoryServiceClient := category.NewCategoryServiceClient(gRCPClientConnection)

	_, clientError := categoryServiceClient.HandleCreate(timeoutBackground, &category.CreateCategoryRequest{
		Name:         categoryCreateDto.Name,
		Description:  categoryCreateDto.Description,
		UserJwtClaim: userJwtClaim,
	})
	exception.ParseGrpcError(ginContext, clientError)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Category has been created", nil))
}

func (categoryHandler *CategoryHandler) Update(ginContext *gin.Context) {
	var updateCategoryDto dto.UpdateCategoryDto
	err := ginContext.ShouldBindBodyWithJSON(&updateCategoryDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	categoryId := ginContext.Param("id")
	timeoutBackground, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	gRCPClientConnection, err := discovery.ServiceConnection(timeoutBackground, "categoryService", categoryHandler.serviceRegistry)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	categoryServiceClient := category.NewCategoryServiceClient(gRCPClientConnection)
	_, clientError := categoryServiceClient.HandleUpdate(timeoutBackground, &category.UpdateCategoryRequest{
		Id:           categoryId,
		Name:         updateCategoryDto.Name,
		Description:  updateCategoryDto.Description,
		UserJwtClaim: userJwtClaim,
	},
	)
	exception.ParseGrpcError(ginContext, clientError)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Category has been updated", nil))
}

func (categoryHandler *CategoryHandler) Delete(ginContext *gin.Context) {
	categoryId := ginContext.Param("id")
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	timeoutBackground, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	gRCPClientConnection, err := discovery.ServiceConnection(timeoutBackground, "categoryService", categoryHandler.serviceRegistry)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	categoryServiceClient := category.NewCategoryServiceClient(gRCPClientConnection)
	_, clientError := categoryServiceClient.HandleDelete(timeoutBackground, &category.DeleteCategoryRequest{
		CategoryId:   categoryId,
		UserJwtClaim: userJwtClaim,
	})
	exception.ParseGrpcError(ginContext, clientError)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Category has been deleted", nil))
}
