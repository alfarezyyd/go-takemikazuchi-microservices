package service

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/category/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/category"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto/user"
)

type CategoryService interface {
	FindAll() *category.QueryCategoryResponses
	HandleCreate(ctx context.Context, userJwtClaim *userDto.JwtClaimDto, categoryCreateDto *dto.CreateCategoryDto)
	HandleUpdate(categoryId string, userJwtClaim *userDto.JwtClaimDto, updateCategoryDto *dto.UpdateCategoryDto) *exception.ClientError
	HandleDelete(categoryId string, userJwtClaim *userDto.JwtClaimDto) *exception.ClientError
}
