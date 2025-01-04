package user

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go-takemikazuchi-api/config"
	"go-takemikazuchi-api/exception"
	"go-takemikazuchi-api/helper"
	"go-takemikazuchi-api/mapper"
	"go-takemikazuchi-api/model"
	"go-takemikazuchi-api/user/dto"
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
	mailerService     *config.MailerService
	identityProvider  *config.IdentityProvider
	viperConfig       *viper.Viper
}

func NewService(
	userRepository Repository,
	dbConnection *gorm.DB,
	validatorInstance *validator.Validate,
	engTranslator ut.Translator,
	mailerService *config.MailerService,
	identityProvider *config.IdentityProvider,
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

func (serviceImpl *ServiceImpl) HandleRegister(createUserDto *dto.CreateUserDto) {
	err := serviceImpl.validatorInstance.Struct(createUserDto)
	exception.ParseValidationError(err, serviceImpl.engTranslator)

	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		userModel := mapper.MapUserDtoIntoUserModel(createUserDto)
		err = gormTransaction.Create(userModel).Error
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest))
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (serviceImpl *ServiceImpl) HandleGenerateOneTimePassword(generateOneTimePassDto *dto.GenerateOtpDto) {
	err := serviceImpl.validatorInstance.Struct(generateOneTimePassDto)
	exception.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		var oneTimePasswordToken model.OneTimePasswordToken
		err = gormTransaction.Where("email = ?", generateOneTimePassDto.Email).First(&userModel).Error
		generatedOneTimePasswordToken := helper.GenerateOneTimePasswordToken()
		hashedGeneratedOneTimePasswordToken, err := bcrypt.GenerateFromPassword([]byte(generatedOneTimePasswordToken), 10)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		oneTimePasswordToken.UserId = userModel.ID
		oneTimePasswordToken.HashedToken = string(hashedGeneratedOneTimePasswordToken)
		oneTimePasswordToken.ExpiresAt = time.Now().Add(15 * time.Minute)
		err = gormTransaction.Create(&oneTimePasswordToken).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		emailPayload := config.EmailPayload{
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
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest))
		return nil
	})
}

func (serviceImpl *ServiceImpl) HandleVerifyOneTimePassword(verifyOtpDto *dto.VerifyOtpDto) {
	err := serviceImpl.validatorInstance.Struct(verifyOtpDto)
	exception.ParseValidationError(err, serviceImpl.engTranslator)
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		var oneTimePasswordToken model.OneTimePasswordToken
		err = gormTransaction.Where("email = ?", verifyOtpDto.Email).First(&userModel).Error
		gormTransaction.Where("user_id = ?", userModel.ID).Order("expires_at desc").First(&oneTimePasswordToken)
		if !(time.Now().Before(oneTimePasswordToken.ExpiresAt)) {
			exception.ThrowClientError(exception.NewClientError(http.StatusBadRequest, "OTP has expired"))
		}
		fmt.Println(oneTimePasswordToken.HashedToken, verifyOtpDto.OneTimePasswordToken)
		err := bcrypt.CompareHashAndPassword([]byte(oneTimePasswordToken.HashedToken), []byte(verifyOtpDto.OneTimePasswordToken))
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest))
		return nil
	})
}

func (serviceImpl *ServiceImpl) HandleGoogleAuthentication() string {
	return serviceImpl.identityProvider.GoogleProviderConfig.AuthCodeURL("randomstate")
}

func (serviceImpl *ServiceImpl) HandleGoogleCallback(tokenState string, queryCode string) *exception.ClientError {
	if tokenState != "randomstate" {
		return &exception.ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    exception.ErrBadRequest,
		}
	}

	googleProviderConfig := serviceImpl.identityProvider.GoogleProviderConfig

	token, err := googleProviderConfig.Exchange(context.Background(), queryCode)
	if err != nil {
		return &exception.ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    exception.ErrBadRequest,
		}
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return &exception.ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    exception.ErrBadRequest,
		}
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return &exception.ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    exception.ErrBadRequest,
		}
	}
	return nil
}

func (serviceImpl *ServiceImpl) HandleLogin(loginUserDto *dto.LoginUserDto) string {
	err := serviceImpl.validatorInstance.Struct(loginUserDto)
	exception.ParseValidationError(err, serviceImpl.engTranslator)
	var tokenString string
	err = serviceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		err = gormTransaction.Where("email = ?", loginUserDto.UserIdentifier).Or("phone_number = ?", loginUserDto.UserIdentifier).First(&userModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(loginUserDto.Password))
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, "User credentials invalid"))
		tokenInstance := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":        userModel.Email,
			"phone_number": helper.ParseNullableValue(userModel.PhoneNumber),
			"exp":          time.Now().Add(time.Hour * 72).Unix(),
		})
		tokenString, err = tokenInstance.SignedString([]byte(serviceImpl.viperConfig.GetString("JWT_SECRET")))
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError))
		return nil
	})
	return tokenString
}
