package category

import (
	"errors"
	"fmt"
	"go-takemikazuchi-microservices/internal/category/dto"
	"go-takemikazuchi-microservices/internal/model"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
	validatorFeature "go-takemikazuchi-microservices/internal/validator"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"go-takemikazuchi-microservices/pkg/mapper"
	"gorm.io/gorm"
	"net/http"
)

type ServiceImpl struct {
	categoryRepository Repository
	dbConnection       *gorm.DB
	validatorService   validatorFeature.Service
}

func NewService(
	categoryRepository Repository,
	dbConnection *gorm.DB,
	validatorService validatorFeature.Service,
) *ServiceImpl {
	return &ServiceImpl{
		categoryRepository: categoryRepository,
		dbConnection:       dbConnection,
		validatorService:   validatorService,
	}
}

func (categoryService *ServiceImpl) FindAll() []dto.CategoryResponseDto {
	categoriesModel := categoryService.categoryRepository.FindAll(categoryService.dbConnection)
	categoriesResponseDto := mapper.MapCategoryModelIntoCategoryResponse(categoriesModel)
	fmt.Println(categoriesModel, categoriesResponseDto)
	return categoriesResponseDto
}

func (categoryService *ServiceImpl) HandleCreate(userJwtClaim *userDto.JwtClaimDto, categoryCreateDto *dto.CreateCategoryDto) *exception.ClientError {
	err := categoryService.validatorService.ValidateStruct(categoryCreateDto)
	categoryService.validatorService.ParseValidationError(err)
	err = categoryService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		var categoryModel model.Category
		err = gormTransaction.Where("email = ?", *userJwtClaim.Email).First(&userModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		if userModel.Role != "Admin" {
			return exception.NewClientError(http.StatusUnauthorized, exception.ErrUnauthorized, errors.New("only admin can create a category"))
		}
		mapper.MapCategoryDtoIntoCategoryModel(&categoryModel, categoryCreateDto)
		err = gormTransaction.Create(&categoryModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return nil
}

func (categoryService *ServiceImpl) HandleUpdate(categoryId string, userJwtClaim *userDto.JwtClaimDto, updateCategoryDto *dto.UpdateCategoryDto) *exception.ClientError {
	err := categoryService.validatorService.ValidateStruct(updateCategoryDto)
	categoryService.validatorService.ParseValidationError(err)
	err = categoryService.validatorService.ValidateVar(categoryId, "required,gte=1")
	categoryService.validatorService.ParseValidationError(err)
	err = categoryService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var categoryModel model.Category
		var userModel model.User
		err = gormTransaction.Where("email = ?", *userJwtClaim.Email).First(&userModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		if userModel.Role != "Admin" {
			return exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err)
		}
		err = gormTransaction.
			Where("categories.id = ?", categoryId).
			First(&categoryModel).
			Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		mapper.MapCategoryDtoIntoCategoryModel(&categoryModel, updateCategoryDto)
		err = gormTransaction.Model(&model.Category{}).Where("id = ?", categoryId).Updates(categoryModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return nil
}

func (categoryService *ServiceImpl) HandleDelete(categoryId string, userJwtClaim *userDto.JwtClaimDto) *exception.ClientError {
	err := categoryService.validatorService.ValidateVar(categoryId, "required,number,gte=1")
	categoryService.validatorService.ParseValidationError(err)
	err = categoryService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var categoryModel model.Category
		var userModel model.User
		err = gormTransaction.Where("email = ?", *userJwtClaim.Email).First(&userModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		if userModel.Role != "Admin" {
			return exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err)
		}
		err = gormTransaction.
			Where("categories.id = ?", categoryId).
			Delete(&categoryModel).
			Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		return nil
	})
	return nil
}
