package category

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/category/dto"
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/helper"
	"go-takemikazuchi-api/mapper"
	"go-takemikazuchi-api/model"
	userDto "go-takemikazuchi-api/user/dto"
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

func (serviceImpl *ServiceImpl) HandleCreate(userJwtClaim *userDto.JwtClaimDto, categoryCreateDto *dto.CreateCategoryDto) *exception.ClientError {
	err := serviceImpl.validationInstance.Struct(categoryCreateDto)
	exception.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		var categoryModel model.Category
		err = gormTransaction.Where("email = ?", *userJwtClaim.Email).First(&userModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		if userModel.Role != "Admin" {
			return exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest)
		}
		mapper.MapCategoryDtoIntoCategoryModel(&categoryModel, categoryCreateDto)
		err = gormTransaction.Create(&categoryModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		return nil
	})
	fmt.Println(err)
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return nil
}

func (serviceImpl *ServiceImpl) HandleUpdate(categoryId string, userJwtClaim *userDto.JwtClaimDto, updateCategoryDto *dto.UpdateCategoryDto) *exception.ClientError {
	err := serviceImpl.validationInstance.Struct(updateCategoryDto)
	exception.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.validationInstance.Var(categoryId, "required,gte=1")
	exception.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var categoryModel model.Category
		var userModel model.User
		err = gormTransaction.Where("email = ?", *userJwtClaim.Email).First(&userModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		if userModel.Role != "Admin" {
			return exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest)
		}
		err = gormTransaction.
			Where("categories.id = ?", categoryId).
			First(&categoryModel).
			Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		mapper.MapCategoryDtoIntoCategoryModel(&categoryModel, updateCategoryDto)
		err = gormTransaction.Where("id = ?", categoryId).Updates(updateCategoryDto).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		return nil
	})
	return nil
}

func (serviceImpl *ServiceImpl) HandleDelete(categoryId string, userJwtClaim *userDto.JwtClaimDto) *exception.ClientError {
	err := serviceImpl.validationInstance.Var(categoryId, "required,number,gte=1")
	exception.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var categoryModel model.Category
		var userModel model.User
		err = gormTransaction.Where("email = ?", *userJwtClaim.Email).First(&userModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		if userModel.Role != "Admin" {
			return exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest)
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
