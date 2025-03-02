package review

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-microservices/internal/review/dto"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"net/http"
)

type Handler struct {
	reviewService Service
}

func NewHandler(reviewService Service) *Handler {
	return &Handler{
		reviewService: reviewService,
	}
}

func (reviewHandler *Handler) Create(ginContext *gin.Context) {
	var createReviewDto dto.CreateReviewDto
	err := ginContext.ShouldBindBodyWithJSON(&createReviewDto)
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	reviewHandler.reviewService.Create(userJwtClaim, &createReviewDto)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Success", nil))
}
