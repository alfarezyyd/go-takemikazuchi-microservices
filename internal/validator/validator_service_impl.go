package validator

import (
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/pkg/exception"
)

type ServiceImpl struct {
	validatorInstance *validator.Validate
	engTranslator     universalTranslator.Translator
}

func NewService(
	validatorInstance *validator.Validate,
	engTranslator universalTranslator.Translator) *ServiceImpl {
	return &ServiceImpl{
		validatorInstance: validatorInstance,
		engTranslator:     engTranslator}
}

func (validatorService *ServiceImpl) ValidateStruct(targetValidatorStruct interface{}) {
	err := validatorService.validatorInstance.Struct(targetValidatorStruct)
	exception.ParseValidationError(err, validatorService.engTranslator)
}

func (validatorService *ServiceImpl) ValidateVar(targetValidatorStruct interface{}, validatorTag string) {
	err := validatorService.validatorInstance.Var(targetValidatorStruct, validatorTag)
	exception.ParseValidationError(err, validatorService.engTranslator)
}
