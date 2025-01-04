package config

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Container struct {
	dbConnection       *gorm.DB
	validationInstance *validator.Validate
	engTranslator      ut.Translator
	viperConfig        *viper.Viper
}

type ContainerOption func(*Container)

func NewContainer(dbConnection *gorm.DB, containerOptions ...ContainerOption) *Container {
	containerInstance := &Container{dbConnection: dbConnection}
	for _, option := range containerOptions {
		option(containerInstance)
	}
	return containerInstance
}

func WithValidator(validationInstance *validator.Validate) ContainerOption {
	return func(containerInstance *Container) { containerInstance.validationInstance = validationInstance }
}

func WithTranslator(engTranslator ut.Translator) ContainerOption {
	return func(containerInstance *Container) { containerInstance.engTranslator = engTranslator }
}
