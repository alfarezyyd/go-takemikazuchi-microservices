package job

import (
	"context"
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go-takemikazuchi-api/internal/category"
	jobDto "go-takemikazuchi-api/internal/job/dto"
	jobResourceFeature "go-takemikazuchi-api/internal/job_resource"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/internal/storage"
	userFeature "go-takemikazuchi-api/internal/user"
	userDto "go-takemikazuchi-api/internal/user/dto"
	userAddressFeature "go-takemikazuchi-api/internal/user_address"
	"go-takemikazuchi-api/internal/worker"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"go-takemikazuchi-api/pkg/mapper"
	"googlemaps.github.io/maps"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
	"strconv"
)

type ServiceImpl struct {
	validatorInstance     *validator.Validate
	jobRepository         Repository
	userRepository        userFeature.Repository
	categoryRepository    category.Repository
	dbConnection          *gorm.DB
	engTranslator         ut.Translator
	jobResourceRepository jobResourceFeature.Repository
	fileStorage           storage.FileStorage
	mapsClient            *maps.Client
	userAddressRepository userAddressFeature.Repository
	workerRepository      worker.Repository
}

func NewService(validatorInstance *validator.Validate,
	jobRepository Repository,
	userRepository userFeature.Repository,
	categoryRepository category.Repository,
	jobResourceRepository jobResourceFeature.Repository,
	dbConnection *gorm.DB,
	engTranslator ut.Translator,
	fileStorage storage.FileStorage,
	mapsClient *maps.Client,
	userAddressRepository userAddressFeature.Repository,
	workerRepository worker.Repository,
) *ServiceImpl {

	return &ServiceImpl{
		validatorInstance:     validatorInstance,
		jobRepository:         jobRepository,
		userRepository:        userRepository,
		categoryRepository:    categoryRepository,
		dbConnection:          dbConnection,
		engTranslator:         engTranslator,
		jobResourceRepository: jobResourceRepository,
		fileStorage:           fileStorage,
		mapsClient:            mapsClient,
		userAddressRepository: userAddressRepository,
		workerRepository:      workerRepository}
}

func (jobService *ServiceImpl) HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto *jobDto.CreateJobDto, uploadedFiles []*multipart.FileHeader) *exception.ClientError {
	err := jobService.validatorInstance.Struct(createJobDto)
	exception.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var jobModel model.Job
		var userModel model.User
		var userAddress model.UserAddress
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		if createJobDto.AddressId == nil {
			geoCodingRequest := &maps.GeocodingRequest{
				LatLng: &maps.LatLng{Lat: createJobDto.Latitude, Lng: createJobDto.Longitude},
			}
			reverseGeocodingResponse, err := jobService.mapsClient.ReverseGeocode(context.Background(), geoCodingRequest)
			helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
			mapper.MapReverseGeocodingIntoUserAddresses(&reverseGeocodingResponse[0], &userAddress, userModel.ID, createJobDto.AdditionalInformationAddress)
			jobService.userAddressRepository.Store(gormTransaction, &userAddress)
		} else {
			jobService.userAddressRepository.FindById(gormTransaction, createJobDto.AddressId, &userAddress)
		}
		isCategoryExists := jobService.categoryRepository.IsCategoryExists(createJobDto.CategoryId, gormTransaction)
		if !isCategoryExists {
			exception.ThrowClientError(exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("category not found")))
		}
		mapper.MapJobDtoIntoJobModel(createJobDto, &jobModel)
		jobModel.UserId = userModel.ID
		jobModel.AddressId = userAddress.ID
		jobService.jobRepository.Store(&jobModel, gormTransaction)
		uuidString := uuid.New().String()
		var allFileName []string
		for _, uploadedFile := range uploadedFiles {
			openedFile, _ := uploadedFile.Open()
			driverLicensePath := fmt.Sprintf("%s-%d-%s", uuidString, jobModel.ID, uploadedFile.Filename)
			_, err = jobService.fileStorage.UploadFile(openedFile, driverLicensePath)
			helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("upload file failed")))
			allFileName = append(allFileName, uploadedFile.Filename)
		}
		resourceModel := mapper.MapStringIntoJobResourceModel(jobModel.ID, allFileName)
		jobService.jobResourceRepository.BulkCreate(gormTransaction, resourceModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	return nil
}

func (jobService *ServiceImpl) HandleUpdate(userJwtClaims *userDto.JwtClaimDto, jobId string, updateJobDto *jobDto.UpdateJobDto, uploadedFiles []*multipart.FileHeader) {
	err := jobService.validatorInstance.Struct(updateJobDto)
	exception.ParseValidationError(err, jobService.engTranslator)
	err = jobService.validatorInstance.Var(jobId, "required|gt=1")
	exception.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		parsedJobId, err := strconv.ParseUint(jobId, 10, 64)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("invalid job id")))
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		jobModel, err := jobService.jobRepository.FindVerifyById(gormTransaction, &userModel.Email, &parsedJobId)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		mapper.MapJobDtoIntoJobModel(updateJobDto, jobModel)
		jobService.jobRepository.Update(jobModel, gormTransaction)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (jobService *ServiceImpl) HandleDelete(userJwtClaims *userDto.JwtClaimDto, jobId string) *exception.ClientError {
	err := jobService.validatorInstance.Var(jobId, "required|gte=1")
	exception.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		jobService.jobRepository.Delete(jobId, userModel.ID, gormTransaction)
		return nil
	})
	return nil
}

func (jobService *ServiceImpl) HandleRequestCompleted(userJwtClaims *userDto.JwtClaimDto, jobId *string) {
	err := jobService.validatorInstance.Var(jobId, "required")
	exception.ParseValidationError(err, jobService.engTranslator)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		parsedJobId, err := strconv.ParseUint(*jobId, 10, 64)
		var userModel model.User
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("invalid job id")))
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		_, err = jobService.jobRepository.VerifyJobOwner(gormTransaction, &userModel.Email, &parsedJobId)
		if err != nil {
			_, err := jobService.jobRepository.VerifyJobWorker(gormTransaction, &userModel.Email, &parsedJobId)
			helper.CheckErrorOperation(err, exception.ParseGormError(err))
		}
		jobModel, err := jobService.jobRepository.FindById(gormTransaction, &parsedJobId)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		jobModel.Status = "Done"
		jobService.jobRepository.Update(jobModel, gormTransaction)
		jobService.workerRepository.DynamicUpdate(gormTransaction, "id = ?", map[string]interface{}{
			"revenue": jobModel.Price,
		}, jobModel.WorkerId)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
