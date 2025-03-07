package service

import (
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/category/internal/category/repository"
	categoryDto "github.com/alfarezyyd/go-takemikazuchi-microservices/category/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/category"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
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

func (categoryService *CategoryServiceImpl) FindAll() *category.QueryCategoryResponses {
	categoriesModel := categoryService.categoryRepository.FindAll(categoryService.dbConnection)
	categoriesResponseDto := mapper.MapCategoryModelIntoCategoryResponse(categoriesModel)
	return &categoriesResponseDto
}

func (categoryService *CategoryServiceImpl) HandleCreate(userJwtClaim *userDto.JwtClaimDto, categoryCreateDto *categoryDto.CreateCategoryDto) *exception.ClientError {
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

func (categoryService *CategoryServiceImpl) HandleUpdate(categoryId string, userJwtClaim *userDto.JwtClaimDto, updateCategoryDto *categoryDto.UpdateCategoryDto) *exception.ClientError {
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
