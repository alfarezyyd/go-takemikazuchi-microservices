package validator

import (
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
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

// ValidateStruct - Validasi struct dengan opsi return error atau panic
func (validatorService *ServiceImpl) ValidateStruct(target interface{}) error {
	return validatorService.validatorInstance.Struct(target)
}

// ValidateVar - Validasi single variable dengan opsi return error atau panic
func (validatorService *ServiceImpl) ValidateVar(target interface{}, validatorTags string) error {
	return validatorService.validatorInstance.Var(target, validatorTags)
}

// ParseValidationError - Parsing error validasi ke dalam format yang lebih mudah dibaca
func (validatorService *ServiceImpl) ParseValidationError(validationError error) {
	if validationError != nil {
		parsedMap := make(map[string]interface{})
		for _, fieldError := range validationError.(validator.ValidationErrors) {
			parsedMap[fieldError.Field()] = fieldError.Translate(validatorService.engTranslator)
		}
		panic(exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, validationError, parsedMap))
	}
}
