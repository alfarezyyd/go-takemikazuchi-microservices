package category

import (
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	dto2 "go-takemikazuchi-api/internal/category/dto"
	model2 "go-takemikazuchi-api/internal/model"
	userDto "go-takemikazuchi-api/internal/user/dto"
	exception2 "go-takemikazuchi-api/pkg/exception"
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

func (serviceImpl *ServiceImpl) HandleCreate(userJwtClaim *userDto.JwtClaimDto, categoryCreateDto *dto2.CreateCategoryDto) *exception2.ClientError {
	err := serviceImpl.validationInstance.Struct(categoryCreateDto)
	exception2.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model2.User
		var categoryModel model2.Category
		err = gormTransaction.Where("email = ?", *userJwtClaim.Email).First(&userModel).Error
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		if userModel.Role != "Admin" {
			return exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, errors.New("only admin can create a category"))
		}
		mapper.MapCategoryDtoIntoCategoryModel(&categoryModel, categoryCreateDto)
		err = gormTransaction.Create(&categoryModel).Error
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		return nil
	})
	helper.CheckErrorOperation(err, exception2.ParseGormError(err))
	return nil
}

func (serviceImpl *ServiceImpl) HandleUpdate(categoryId string, userJwtClaim *userDto.JwtClaimDto, updateCategoryDto *dto2.UpdateCategoryDto) *exception2.ClientError {
	err := serviceImpl.validationInstance.Struct(updateCategoryDto)
	exception2.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.validationInstance.Var(categoryId, "required,gte=1")
	exception2.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var categoryModel model2.Category
		var userModel model2.User
		err = gormTransaction.Where("email = ?", *userJwtClaim.Email).First(&userModel).Error
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		if userModel.Role != "Admin" {
			return exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err)
		}
		err = gormTransaction.
			Where("categories.id = ?", categoryId).
			First(&categoryModel).
			Error
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		mapper.MapCategoryDtoIntoCategoryModel(&categoryModel, updateCategoryDto)
		err = gormTransaction.Model(&model2.Category{}).Where("id = ?", categoryId).Updates(categoryModel).Error
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		return nil
	})
	helper.CheckErrorOperation(err, exception2.ParseGormError(err))
	return nil
}

func (serviceImpl *ServiceImpl) HandleDelete(categoryId string, userJwtClaim *userDto.JwtClaimDto) *exception2.ClientError {
	err := serviceImpl.validationInstance.Var(categoryId, "required,number,gte=1")
	exception2.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var categoryModel model2.Category
		var userModel model2.User
		err = gormTransaction.Where("email = ?", *userJwtClaim.Email).First(&userModel).Error
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		if userModel.Role != "Admin" {
			return exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err)
		}
		err = gormTransaction.
			Where("categories.id = ?", categoryId).
			Delete(&categoryModel).
			Error
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		return nil
	})
	return nil
}
