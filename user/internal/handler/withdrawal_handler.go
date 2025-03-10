package handler

import (
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/service"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	withdrawalService service.WithdrawalService
}

func NewHandler(withdrawalService service.WithdrawalService) *Handler {
	return &Handler{
		withdrawalService: withdrawalService,
	}
}

func (withdrawalHandler *Handler) FindAll(ginContext *gin.Context) {
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	withdrawalsModel := withdrawalHandler.withdrawalService.FindAll(userJwtClaim)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Success", withdrawalsModel))
}

func (withdrawalHandler *Handler) Create(ginContext *gin.Context) {
	var createWithdrawalDto userDto.CreateWithdrawalDto
	err := ginContext.ShouldBindBodyWithJSON(&createWithdrawalDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("error parsing body")))
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	withdrawalHandler.withdrawalService.Create(userJwtClaim, &createWithdrawalDto)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Success", nil))
}

func (withdrawalHandler *Handler) Update(ginContext *gin.Context) {
	withdrawalId := ginContext.Param("withdrawalId")
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	withdrawalHandler.withdrawalService.Update(userJwtClaim, &withdrawalId)
	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Success", nil))
}
