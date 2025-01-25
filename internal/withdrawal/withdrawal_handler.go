package withdrawal

import (
	"errors"
	"github.com/gin-gonic/gin"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/internal/withdrawal/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
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
