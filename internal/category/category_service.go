package category

import (
	dto2 "go-takemikazuchi-api/internal/category/dto"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
)

type Service interface {
	HandleCreate(userJwtClaim *userDto.JwtClaimDto, categoryCreateDto *dto2.CreateCategoryDto) *exception.ClientError
	HandleUpdate(categoryId string, userJwtClaim *userDto.JwtClaimDto, updateCategoryDto *dto2.UpdateCategoryDto) *exception.ClientError
	HandleDelete(categoryId string, userJwtClaim *userDto.JwtClaimDto) *exception.ClientError
}
