package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go-takemikazuchi-api/configs"
	model2 "go-takemikazuchi-api/internal/model"
	dto2 "go-takemikazuchi-api/internal/user/dto"
	exception2 "go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"go-takemikazuchi-api/pkg/mapper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"time"
)

type ServiceImpl struct {
	userRepository    Repository
	dbConnection      *gorm.DB
	validatorInstance *validator.Validate
	engTranslator     ut.Translator
	mailerService     *configs.MailerService
	identityProvider  *configs.IdentityProvider
	viperConfig       *viper.Viper
}

func NewService(
	userRepository Repository,
	dbConnection *gorm.DB,
	validatorInstance *validator.Validate,
	engTranslator ut.Translator,
	mailerService *configs.MailerService,
	identityProvider *configs.IdentityProvider,
	viperConfig *viper.Viper) *ServiceImpl {
	return &ServiceImpl{
		userRepository:    userRepository,
		dbConnection:      dbConnection,
		validatorInstance: validatorInstance,
		engTranslator:     engTranslator,
		mailerService:     mailerService,
		identityProvider:  identityProvider,
		viperConfig:       viperConfig,
	}
}

func (serviceImpl *ServiceImpl) HandleRegister(createUserDto *dto2.CreateUserDto) {
	err := serviceImpl.validatorInstance.Struct(createUserDto)
	exception2.ParseValidationError(err, serviceImpl.engTranslator)

	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		userModel := mapper.MapUserDtoIntoUserModel(createUserDto)
		err = gormTransaction.Create(userModel).Error
		helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
		return nil
	})
	helper.CheckErrorOperation(err, exception2.ParseGormError(err))
}

func (serviceImpl *ServiceImpl) HandleGenerateOneTimePassword(generateOneTimePassDto *dto2.GenerateOtpDto) {
	err := serviceImpl.validatorInstance.Struct(generateOneTimePassDto)
	exception2.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model2.User
		var oneTimePasswordToken model2.OneTimePasswordToken
		err = gormTransaction.Where("email = ?", generateOneTimePassDto.Email).First(&userModel).Error
		generatedOneTimePasswordToken := helper.GenerateOneTimePasswordToken()
		hashedGeneratedOneTimePasswordToken, err := bcrypt.GenerateFromPassword([]byte(generatedOneTimePasswordToken), 10)
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		oneTimePasswordToken.UserId = userModel.ID
		oneTimePasswordToken.HashedToken = string(hashedGeneratedOneTimePasswordToken)
		oneTimePasswordToken.ExpiresAt = time.Now().Add(15 * time.Minute)
		err = gormTransaction.Create(&oneTimePasswordToken).Error
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		emailPayload := configs.EmailPayload{
			Title:     "One Time Verification Token",
			Recipient: generateOneTimePassDto.Email,
			Body:      fmt.Sprintf("One Time Password %s", generatedOneTimePasswordToken),
			Sender:    serviceImpl.mailerService.ViperConfig.GetString(""),
		}

		projectRoot, _ := os.Getwd() // Mendapatkan root path proyek
		templateFile := fmt.Sprintf("%s/public/static/email_template.html", projectRoot)
		err = serviceImpl.mailerService.SendEmail(
			generateOneTimePassDto.Email,
			"One Time Verification Token",
			templateFile,
			emailPayload)
		helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
		return nil
	})
}

func (serviceImpl *ServiceImpl) HandleVerifyOneTimePassword(verifyOtpDto *dto2.VerifyOtpDto) {
	err := serviceImpl.validatorInstance.Struct(verifyOtpDto)
	exception2.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model2.User
		var oneTimePasswordToken model2.OneTimePasswordToken
		err = gormTransaction.Where("email = ?", verifyOtpDto.Email).First(&userModel).Error
		gormTransaction.Where("user_id = ?", userModel.ID).Order("expires_at desc").First(&oneTimePasswordToken)
		if !(time.Now().Before(oneTimePasswordToken.ExpiresAt)) {
			exception2.ThrowClientError(exception2.NewClientError(http.StatusBadRequest, "OTP has expired", err))
		}
		err := bcrypt.CompareHashAndPassword([]byte(oneTimePasswordToken.HashedToken), []byte(verifyOtpDto.OneTimePasswordToken))
		helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err))
		return nil
	})
}

func (serviceImpl *ServiceImpl) HandleGoogleAuthentication() string {
	return serviceImpl.identityProvider.GoogleProviderConfig.AuthCodeURL("randomstate")
}

func (serviceImpl *ServiceImpl) HandleGoogleCallback(tokenState string, queryCode string) *exception2.ClientError {
	if tokenState != "randomstate" {
		return exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, errors.New("invalid token state"))
	}

	googleProviderConfig := serviceImpl.identityProvider.GoogleProviderConfig

	token, err := googleProviderConfig.Exchange(context.Background(), queryCode)
	if err != nil {
		return exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err)
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err)
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return exception2.NewClientError(http.StatusBadRequest, exception2.ErrBadRequest, err)
	}
	return nil
}

func (serviceImpl *ServiceImpl) HandleLogin(loginUserDto *dto2.LoginUserDto) string {
	err := serviceImpl.validatorInstance.Struct(loginUserDto)
	exception2.ParseValidationError(err, serviceImpl.engTranslator)
	var tokenString string
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model2.User
		err = gormTransaction.Where("email = ?", loginUserDto.UserIdentifier).Or("phone_number = ?", loginUserDto.UserIdentifier).First(&userModel).Error
		helper.CheckErrorOperation(err, exception2.ParseGormError(err))
		err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(loginUserDto.Password))
		helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusBadRequest, "User credentials invalid", err))
		tokenInstance := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":        userModel.Email,
			"phone_number": helper.ParseNullableValue(userModel.PhoneNumber),
			"exp":          time.Now().Add(time.Hour * 72).Unix(),
		})
		tokenString, err = tokenInstance.SignedString([]byte(serviceImpl.viperConfig.GetString("JWT_SECRET")))
		helper.CheckErrorOperation(err, exception2.NewClientError(http.StatusInternalServerError, exception2.ErrInternalServerError, err))
		return nil
	})
	return tokenString
}
