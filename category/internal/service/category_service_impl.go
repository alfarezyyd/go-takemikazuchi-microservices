package service

import (
	"errors"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-category/internal/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-category/internal/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-common/pkg/mapper"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices-common/pkg/validator"
	"gorm.io/gorm"
	"net/http"
)

type CategoryServiceImpl struct {
	categoryRepository repository.Repository
	dbConnection       *gorm.DB
	validatorService   validatorFeature.Service
}

func NewService(
	categoryRepository repository.Repository,
	dbConnection *gorm.DB,
	validatorService validatorFeature.Service,
) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		categoryRepository: categoryRepository,
		dbConnection:       dbConnection,
		validatorService:   validatorService,
	}
}

func (categoryService *CategoryServiceImpl) FindAll() []dto.CategoryResponseDto {
	categoriesModel := categoryService.categoryRepository.FindAll(categoryService.dbConnection)
	categoriesResponseDto := mapper.MapCategoryModelIntoCategoryResponse(categoriesModel)
	fmt.Println(categoriesModel, categoriesResponseDto)
	return categoriesResponseDto
}

func (categoryService *CategoryServiceImpl) HandleCreate(userJwtClaim *userDto.JwtClaimDto, categoryCreateDto *dto.CreateCategoryDto) *exception.ClientError {
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

func (categoryService *CategoryServiceImpl) HandleUpdate(categoryId string, userJwtClaim *userDto.JwtClaimDto, updateCategoryDto *dto.UpdateCategoryDto) *exception.ClientError {
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

func (categoryService *CategoryServiceImpl) HandleDelete(categoryId string, userJwtClaim *userDto.JwtClaimDto) *exception.ClientError {
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
