package review

import (
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/review/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices-user/internal/user/dto"
	"github.com/gin-gonic/gin"
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
