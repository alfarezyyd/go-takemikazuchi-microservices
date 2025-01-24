package review

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-takemikazuchi-api/internal/review/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
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
	fmt.Println(err)
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
	reviewHandler.reviewService.Create(userJwtClaim, &createReviewDto)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Success", nil))
}
