package category

import (
	"go-takemikazuchi-microservices/internal/category/dto"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
	"go-takemikazuchi-microservices/pkg/exception"
)

type Service interface {
	FindAll() []dto.CategoryResponseDto
	HandleCreate(userJwtClaim *userDto.JwtClaimDto, categoryCreateDto *dto.CreateCategoryDto) *exception.ClientError
	HandleUpdate(categoryId string, userJwtClaim *userDto.JwtClaimDto, updateCategoryDto *dto.UpdateCategoryDto) *exception.ClientError
	HandleDelete(categoryId string, userJwtClaim *userDto.JwtClaimDto) *exception.ClientError
}
