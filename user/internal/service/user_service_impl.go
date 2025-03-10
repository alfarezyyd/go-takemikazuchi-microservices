package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"time"
)

type UserServiceImpl struct {
	userRepository   repository.UserRepository
	dbConnection     *gorm.DB
	mailerService    *configs.MailerService
	identityProvider *configs.IdentityProvider
	viperConfig      *viper.Viper
	validatorService validatorFeature.Service
}

func NewUserService(
	validatorService validatorFeature.Service,
	userRepository repository.UserRepository,
	dbConnection *gorm.DB,
	mailerService *configs.MailerService,
	identityProvider *configs.IdentityProvider,
	viperConfig *viper.Viper) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository:   userRepository,
		dbConnection:     dbConnection,
		mailerService:    mailerService,
		identityProvider: identityProvider,
		viperConfig:      viperConfig,
		validatorService: validatorService,
	}
}

func (userService *UserServiceImpl) HandleRegister(createUserDto *dto.CreateUserDto) error {
	err := userService.validatorService.ValidateStruct(createUserDto)
	userService.validatorService.ParseValidationError(err)
	err = userService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		isUserExists, err := userService.userRepository.IsUserExists(gormTransaction, "phone_number = ? OR email = ?", createUserDto.PhoneNumber, createUserDto.Email)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		if isUserExists {
			exception.ThrowClientError(exception.NewClientError(http.StatusBadRequest, "Email or phone number has been registered", errors.New("duplicate email")))
		}
		userModel := mapper.MapUserDtoIntoUserModel(createUserDto)
		err = gormTransaction.Create(userModel).Error
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
		if createUserDto.Email != "" {
			userService.HandleGenerateOneTimePassword(&dto.GenerateOtpDto{
				Email:  createUserDto.Email,
				UserId: userModel.ID,
			}, gormTransaction)
		}
		return nil
	})
	gormError := exception.ParseGormError(err)
	fmt.Println(gormError)
	if gormError != nil {
		return exception.ParseIntoGrpcError(gormError)
	}
	return nil
}

func (userService *UserServiceImpl) HandleGenerateOneTimePassword(generateOneTimePassDto *dto.GenerateOtpDto, externalGormTransaction *gorm.DB) {
	err := userService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		if externalGormTransaction != nil {
			gormTransaction = externalGormTransaction
		}
		var userModel model.User
		var oneTimePasswordToken model.OneTimePasswordToken
		fmt.Println(generateOneTimePassDto)
		err := gormTransaction.Where("email = ?", generateOneTimePassDto.Email).First(&userModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		generatedOneTimePasswordToken := helper.GenerateOneTimePasswordToken()
		hashedGeneratedOneTimePasswordToken, err := bcrypt.GenerateFromPassword([]byte(generatedOneTimePasswordToken), 10)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		oneTimePasswordToken.UserId = userModel.ID
		oneTimePasswordToken.HashedToken = string(hashedGeneratedOneTimePasswordToken)
		oneTimePasswordToken.ExpiresAt = time.Now().Add(15 * time.Minute)
		err = gormTransaction.Create(&oneTimePasswordToken).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		emailPayload := configs.EmailPayload{
			Title:     "One Time Verification Token",
			Recipient: generateOneTimePassDto.Email,
			Body:      fmt.Sprintf("One Time Password %s", generatedOneTimePasswordToken),
			Sender:    userService.mailerService.ViperConfig.GetString(""),
		}

		projectRoot, _ := os.Getwd() // Mendapatkan root path proyek
		templateFile := fmt.Sprintf("%s/public/static/email_template.html", projectRoot)
		err = userService.mailerService.SendEmail(
			generateOneTimePassDto.Email,
			"One Time Verification Token",
			templateFile,
			emailPayload)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (userService *UserServiceImpl) HandleVerifyOneTimePassword(verifyOtpDto *dto.VerifyOtpDto) {
	err := userService.validatorService.ValidateStruct(verifyOtpDto)
	userService.validatorService.ParseValidationError(err)
	err = userService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		var oneTimePasswordToken model.OneTimePasswordToken
		err = gormTransaction.Where("email = ?", verifyOtpDto.Email).First(&userModel).Error
		gormTransaction.Where("user_id = ?", userModel.ID).Order("expires_at desc").First(&oneTimePasswordToken)
		if !(time.Now().Before(oneTimePasswordToken.ExpiresAt)) {
			exception.ThrowClientError(exception.NewClientError(http.StatusBadRequest, "OTP has expired", err))
		}
		err := bcrypt.CompareHashAndPassword([]byte(oneTimePasswordToken.HashedToken), []byte(verifyOtpDto.OneTimePasswordToken))
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err))
		return nil
	})
}

func (userService *UserServiceImpl) HandleGoogleAuthentication() string {
	return userService.identityProvider.GoogleProviderConfig.AuthCodeURL("randomstate")
}

func (userService *UserServiceImpl) HandleGoogleCallback(tokenState string, queryCode string) *exception.ClientError {
	if tokenState != "randomstate" {
		return exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("invalid token state"))
	}

	googleProviderConfig := userService.identityProvider.GoogleProviderConfig

	token, err := googleProviderConfig.Exchange(context.Background(), queryCode)
	if err != nil {
		return exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err)
	}

	resp, err := http.Get("https://www.googlemicroservicess.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err)
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, err)
	}
	return nil
}

func (userService *UserServiceImpl) HandleLogin(loginUserDto *user.LoginUserRequest) string {
	err := userService.validatorService.ValidateStruct(loginUserDto)
	userService.validatorService.ParseValidationError(err)
	var tokenString string
	err = userService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		err = gormTransaction.Where("email = ?", loginUserDto.UserIdentifier).Or("phone_number = ?", loginUserDto.UserIdentifier).First(&userModel).Error
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(loginUserDto.Password))
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, "User credentials invalid", err))
		tokenInstance := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":        userModel.Email,
			"phone_number": helper.ParseNullableValue(userModel.PhoneNumber),
			"exp":          time.Now().Add(time.Hour * 72).Unix(),
		})
		tokenString, err = tokenInstance.SignedString([]byte(userService.viperConfig.GetString("JWT_SECRET")))
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
		return nil
	})
	return tokenString
}

func (userService *UserServiceImpl) FindByIdentifier(userIdentifierDto *dto.UserIdentifierDto) *dto.UserResponseDto {
	err := userService.validatorService.ValidateStruct(userIdentifierDto)
	userService.validatorService.ParseValidationError(err)

	var userModel model.User
	err = userService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		err := userService.dbConnection.Where(
			"email = ? OR phone_number = ?",
			userIdentifierDto.Email,
			userIdentifierDto.PhoneNumber).First(&userModel).Error
		fmt.Println(err)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return mapper.MapUserModelIntoUserResponse(&userModel)
}
