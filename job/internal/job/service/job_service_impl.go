package service

import (
	"errors"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/storage"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/job/repository"
	jobDto "github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto/job"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto/user"
	"github.com/google/uuid"
	"googlemaps.github.io/maps"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
)

type JobServiceImpl struct {
	validatorService validatorFeature.Service
	jobRepository    repository.JobRepository
	//userRepository        userFeature.Repository
	//categoryRepository    category.Repository
	dbConnection *gorm.DB
	//jobResourceRepository jobResourceFeature.Repository
	fileStorage storage.FileStorage
	mapsClient  *maps.Client
	//userAddressRepository userAddressFeature.Repository
	//workerRepository      worker.Repository
	serviceDiscovery discovery.ServiceRegistry
}

func NewJobService(
	jobRepository repository.JobRepository,
	//userRepository userFeature.Repository,
	//categoryRepository category.Repository,
	//jobResourceRepository jobResourceFeature.Repository,
	dbConnection *gorm.DB,
	fileStorage storage.FileStorage,
	mapsClient *maps.Client,
	//userAddressRepository userAddressFeature.Repository,
	//workerRepository worker.Repository,
	validatorService validatorFeature.Service,
	serviceDiscovery discovery.ServiceRegistry,
) *JobServiceImpl {
	return &JobServiceImpl{
		jobRepository: jobRepository,
		//userRepository:        userRepository,
		//categoryRepository:    categoryRepository,
		dbConnection: dbConnection,
		//jobResourceRepository: jobResourceRepository,
		fileStorage: fileStorage,
		mapsClient:  mapsClient,
		//userAddressRepository: userAddressRepository,
		validatorService: validatorService,
		//workerRepository:      workerRepository
		serviceDiscovery: serviceDiscovery,
	}
}

func (jobService *JobServiceImpl) HandleCreate(userJwtClaims *userDto.JwtClaimDto, createJobDto *jobDto.CreateJobDto, uploadedFiles []*multipart.FileHeader) *exception.ClientError {
	err := jobService.validatorService.ValidateStruct(createJobDto)
	jobService.validatorService.ParseValidationError(err)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var jobModel model.Job
		var userModel model.User
		var userAddress model.UserAddress
		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		if createJobDto.AddressId == nil {
			//geoCodingRequest := &maps.GeocodingRequest{
			//	LatLng: &maps.LatLng{Lat: createJobDto.Latitude, Lng: createJobDto.Longitude},
			//}
			//reverseGeocodingResponse, err := jobService.mapsClient.ReverseGeocode(context.Background(), geoCodingRequest)
			//helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("bad request")))
			//mapper.MapReverseGeocodingIntoUserAddresses(&reverseGeocodingResponse[0], &userAddress, userModel.ID, createJobDto.AdditionalInformationAddress)
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
		//resourceModel := mapper.MapStringIntoJobResourceModel(jobModel.ID, allFileName)
		//jobService.jobResourceRepository.BulkCreate(gormTransaction, resourceModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	return nil
}

//func (jobService *JobServiceImpl) HandleUpdate(userJwtClaims *userDto.JwtClaimDto, jobId string, updateJobDto *jobDto.UpdateJobDto, uploadedFiles []*multipart.FileHeader) {
//	err := jobService.validatorService.ValidateStruct(updateJobDto)
//	jobService.validatorService.ParseValidationError(err)
//	err = jobService.validatorService.ValidateVar(jobId, "required|gt=1")
//	jobService.validatorService.ParseValidationError(err)
//	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
//		var userModel model.User
//		parsedJobId, err := strconv.ParseUint(jobId, 10, 64)
//		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("invalid job id")))
//		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
//		jobModel, err := jobService.jobRepository.FindVerifyById(gormTransaction, &userModel.Email, &parsedJobId)
//		if jobModel.CategoryId != updateJobDto.CategoryId {
//			isCategoryExists := jobService.categoryRepository.IsCategoryExists(updateJobDto.CategoryId, gormTransaction)
//			if !isCategoryExists {
//				exception.ThrowClientError(exception.NewClientError(http.StatusNotFound, exception.ErrNotFound, errors.New("category not found")))
//			}
//		}
//		helper.CheckErrorOperation(err, exception.ParseGormError(err))
//		mapper.MapJobDtoIntoJobModel(updateJobDto, jobModel)
//		jobService.jobRepository.Update(jobModel, gormTransaction)
//		resourceModel := jobService.UpdateUploadedFiles(uploadedFiles, jobModel.ID)
//		if resourceModel != nil {
//			jobService.jobResourceRepository.BulkCreate(gormTransaction, resourceModel)
//		}
//		if len(updateJobDto.DeletedFilesName) != 0 {
//			countFile := jobService.jobResourceRepository.CountBulkByName(gormTransaction, jobModel.ID, updateJobDto.DeletedFilesName)
//			jobService.DeleteRequestedFile(updateJobDto.DeletedFilesName, countFile)
//			jobService.jobResourceRepository.DeleteBulkByName(gormTransaction, jobModel.ID, updateJobDto.DeletedFilesName)
//		}
//		return nil
//	})
//	fmt.Print(err)
//	helper.CheckErrorOperation(err, exception.ParseGormError(err))
//}
//
//func (jobService *JobServiceImpl) HandleDelete(userJwtClaims *userDto.JwtClaimDto, jobId string) *exception.ClientError {
//	err := jobService.validatorService.ValidateVar(jobId, "required|gte=1")
//	jobService.validatorService.ParseValidationError(err)
//	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
//		var userModel model.User
//		parsedJobId, err := strconv.ParseUint(jobId, 10, 64)
//		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("invalid job id")))
//		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
//		_, err = jobService.jobRepository.VerifyJobOwner(gormTransaction, userJwtClaims.Email, &parsedJobId)
//		if err != nil {
//			exception.ThrowClientError(exception.NewClientError(http.StatusUnauthorized, exception.ErrUnauthorized, errors.New("job not belong to user")))
//		}
//		jobService.jobResourceRepository.DeleteBulkByJobId(gormTransaction, &parsedJobId)
//		jobService.jobRepository.Delete(jobId, userModel.ID, gormTransaction)
//		return nil
//	})
//	return nil
//}
//
//func (jobService *JobServiceImpl) HandleRequestCompleted(userJwtClaims *userDto.JwtClaimDto, jobId *string) {
//	err := jobService.validatorService.ValidateVar(jobId, "required")
//	jobService.validatorService.ParseValidationError(err)
//	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
//		parsedJobId, err := strconv.ParseUint(*jobId, 10, 64)
//		var userModel model.User
//		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("invalid job id")))
//		jobService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
//		_, err = jobService.jobRepository.VerifyJobOwner(gormTransaction, &userModel.Email, &parsedJobId)
//		if err != nil {
//			_, err := jobService.jobRepository.VerifyJobWorker(gormTransaction, &userModel.Email, &parsedJobId)
//			helper.CheckErrorOperation(err, exception.ParseGormError(err))
//		}
//		jobModel, err := jobService.jobRepository.FindById(gormTransaction, &parsedJobId)
//		helper.CheckErrorOperation(err, exception.ParseGormError(err))
//		jobModel.Status = "Done"
//		jobService.jobRepository.Update(jobModel, gormTransaction)
//		jobService.workerRepository.DynamicUpdate(gormTransaction, "id = ?", map[string]interface{}{
//			"revenue": jobModel.Price,
//		}, jobModel.WorkerId)
//		return nil
//	})
//	helper.CheckErrorOperation(err, exception.ParseGormError(err))
//}
//
//func (jobService *JobServiceImpl) UpdateUploadedFiles(uploadedFiles []*multipart.FileHeader, jobId uint64) []*model.JobResource {
//	var resourceModel []*model.JobResource
//	uuidString := uuid.New().String()
//	if len(uploadedFiles) != 0 {
//		var allFileName []string
//		for _, uploadedFile := range uploadedFiles {
//			openedFile, _ := uploadedFile.Open()
//			_, err := jobService.fileStorage.UploadFile(openedFile, fmt.Sprintf("%s-%s", uuidString, uploadedFile.Filename))
//			helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("upload file failed")))
//			allFileName = append(allFileName, fmt.Sprintf("%s-%s", uuidString, uploadedFile.Filename))
//		}
//		resourceModel = mapper.MapStringIntoJobResourceModel(jobId, allFileName)
//	}
//	return resourceModel
//}
//
//func (jobService *JobServiceImpl) DeleteRequestedFile(deletedFilesName []string, countFile int) {
//	if countFile != len(deletedFilesName) {
//		exception.ThrowClientError(exception.NewClientError(http.StatusNotFound, "Some files not found", errors.New("count file not equal")))
//	}
//	for _, deletedFileName := range deletedFilesName {
//		_ = jobService.fileStorage.DeleteFile(deletedFileName)
//	}
//}
