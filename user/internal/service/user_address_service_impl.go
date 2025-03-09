package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
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

func (userAddressServiceImpl *UserAddressServiceImpl) Create(ctx context.Context, createUserAddressDto *dto.CreateUserAddressDto) uint64 {
	reverseResponse, err := userAddressServiceImpl.nominatimHttpClient.HTTPClient.Get(fmt.Sprintf("%s/reverse?lat=%f&lon=%f&format=json", *userAddressServiceImpl.nominatimHttpClient.BaseURL, createUserAddressDto.Latitude, createUserAddressDto.Longitude))
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("failed call reverse geocoding")))
	responseBody, readErr := io.ReadAll(reverseResponse.Body)
	helper.CheckErrorOperation(readErr, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("failed call reverse geocoding")))
	var userLocation dto.UserLocation
	jsonErr := json.Unmarshal(responseBody, &userLocation)
	helper.CheckErrorOperation(jsonErr, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("failed call reverse geocoding")))
	userAddress := mapper.MapLocationToUserAddress(userLocation, createUserAddressDto.UserId)
	err = userAddressServiceImpl.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		userAddressServiceImpl.userAddressRepository.Store(gormTransaction, &userAddress)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return userAddress.ID
}

func (userAddressServiceImpl *UserAddressServiceImpl) FindById(ctx context.Context, searchUserAddressDto *dto.SearchUserAddressDto) {

}
