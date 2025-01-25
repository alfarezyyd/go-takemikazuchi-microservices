package category

import (
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/internal/category/dto"
	"go-takemikazuchi-api/internal/model"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"go-takemikazuchi-api/pkg/mapper"
	"gorm.io/gorm"
	"net/http"
)

type ServiceImpl struct {
	categoryRepository Repository
	dbConnection       *gorm.DB
	validationInstance *validator.Validate
	engTranslator      ut.Translator
}

func NewService(
	categoryRepository Repository,
	dbConnection *gorm.DB,
	validatorInstance *validator.Validate,
	engTranslator ut.Translator) *ServiceImpl {
	return &ServiceImpl{
		categoryRepository: categoryRepository,
		dbConnection:       dbConnection,
		validationInstance: validatorInstance,
		engTranslator:      engTranslator,
	}
}

func (categoryService *ServiceImpl) FindAll() []dto.CategoryResponseDto {
	categoriesModel := categoryService.categoryRepository.FindAll(categoryService.dbConnection)
	categoriesResponseDto := mapper.MapCategoryModelIntoCategoryResponse(categoriesModel)
	fmt.Println(categoriesModel, categoriesResponseDto)
	return categoriesResponseDto
}

func (categoryService *ServiceImpl) HandleCreate(userJwtClaim *userDto.JwtClaimDto, categoryCreateDto *dto.CreateCategoryDto) *exception.ClientError {
	err := categoryService.validationInstance.Struct(categoryCreateDto)
	exception.ParseValidationError(err, categoryService.engTranslator)
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
	err := categoryService.validationInstance.Struct(updateCategoryDto)
	exception.ParseValidationError(err, categoryService.engTranslator)
	err = categoryService.validationInstance.Var(categoryId, "required,gte=1")
	exception.ParseValidationError(err, categoryService.engTranslator)
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
	err := categoryService.validationInstance.Var(categoryId, "required,number,gte=1")
	exception.ParseValidationError(err, categoryService.engTranslator)
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
