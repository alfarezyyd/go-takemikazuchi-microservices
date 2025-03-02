package withdrawal

import (
	"errors"
	"github.com/gin-gonic/gin"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
	"go-takemikazuchi-microservices/internal/withdrawal/dto"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"net/http"
)

type Handler struct {
	withdrawalService Service
}

func NewHandler(withdrawalService Service) *Handler {
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
	var createWithdrawalDto dto.CreateWithdrawalDto
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
