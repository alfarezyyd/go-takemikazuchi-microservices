package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type UserAddressServiceImpl struct {
	userRepository        repository.UserRepository
	dbConnection          *gorm.DB
	userAddressRepository repository.UserAddressRepository
	validatorService      validatorFeature.Service
	nominatimHttpClient   *configs.HttpClient
}

func NewUserAddressServiceImpl(userRepository repository.UserRepository, dbConnection *gorm.DB, userAddressRepository repository.UserAddressRepository, validatorService validatorFeature.Service, nominatimHttpClient *configs.HttpClient,
) *UserAddressServiceImpl {
	return &UserAddressServiceImpl{
		userRepository:        userRepository,
		dbConnection:          dbConnection,
		userAddressRepository: userAddressRepository,
		validatorService:      validatorService,
		nominatimHttpClient:   nominatimHttpClient,
	}
}

func (userAddressServiceImpl *UserAddressServiceImpl) Create(ctx context.Context, createUserAddressDto *dto.CreateUserAddressDto) {
	reverseResponse, err := userAddressServiceImpl.nominatimHttpClient.HTTPClient.Get(fmt.Sprintf("%s/reverse?lat=%f&lon=%f&format=json", *userAddressServiceImpl.nominatimHttpClient.BaseURL, createUserAddressDto.Latitude, createUserAddressDto.Longitude))
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("failed call reverse geocoding")))
	responseBody, readErr := io.ReadAll(reverseResponse.Body)
	helper.CheckErrorOperation(readErr, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("failed call reverse geocoding")))
	fmt.Println(string(responseBody))
}

func (userAddressServiceImpl *UserAddressServiceImpl) FindById(ctx context.Context, searchUserAddressDto *dto.SearchUserAddressDto) {

}
