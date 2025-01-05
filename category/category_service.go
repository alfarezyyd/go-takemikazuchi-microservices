package category

import (
	"go-takemikazuchi-api/category/dto"
	"go-takemikazuchi-api/exception"
	userDto "go-takemikazuchi-api/user/dto"
)

type Service interface {
	HandleCreate(userJwtClaim *userDto.JwtClaimDto, categoryCreateDto *dto.CreateCategoryDto) *exception.ClientError
	HandleUpdate(categoryId string, userJwtClaim *userDto.JwtClaimDto, updateCategoryDto *dto.UpdateCategoryDto) *exception.ClientError
	HandleDelete(categoryId string, userJwtClaim *userDto.JwtClaimDto) *exception.ClientError
}
