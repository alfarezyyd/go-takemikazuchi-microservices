package validator

type Service interface {
	ValidateStruct(targetValidatorStruct interface{})
	ValidateVar(targetValidatorStruct interface{}, validatorTag string)
}
