package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/configs"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/category"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	userAddressGrpc "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user_address"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/storage"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/repository"
	jobDto "github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"googlemaps.github.io/maps"
	"gorm.io/gorm"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
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
	rabbitMQ         *configs.RabbitMQConsumer
}

func NewJobService(
	jobRepository repository.JobRepository,
	//userRepository userFeature.Repository,
	//categoryRepository category.Repository,
	//jobResourceRepository jobResourceFeature.Repository,
	dbConnection *gorm.DB,
	fileStorage storage.FileStorage,
	//userAddressRepository userAddressFeature.Repository,
	//workerRepository worker.Repository,
	validatorService validatorFeature.Service,
	serviceDiscovery discovery.ServiceRegistry,
	rabbitMQ *configs.RabbitMQConsumer,
) *JobServiceImpl {

	go rabbitMQ.StartListening(func(body []byte) error {
		var eventData map[string]interface{}
		if err := json.Unmarshal(body, &eventData); err != nil {
			return err
		}

		// 3. Proses pesan yang diterima
		log.Printf("Processing order update: %v", eventData)

		// Contoh: Update status job di database
		orderID, ok := eventData["order_id"].(string)
		if !ok {
			return fmt.Errorf("invalid order ID format")
		}
		status, ok := eventData["status"].(string)
		if !ok {
			return fmt.Errorf("invalid status format")
		}
		jobId, ok := eventData["job_id"].(string)
		if !ok {
			return fmt.Errorf("invalid job id format")
		}

		err := dbConnection.Transaction(func(tx *gorm.DB) error {
			// 4. Update Job Berdasarkan Order
			parsedJobId, _ := strconv.ParseUint(jobId, 10, 64)
			job, err := jobRepository.FindById(tx, &parsedJobId)
			if err != nil {
				return fmt.Errorf("job not found for order ID: %s", orderID)
			}
			switch status {
			case "capture":
				job.Status = "On Working"
				break
			case "settlement":
				job.Status = "On Working"
				break
			case "deny":
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
				break
			case "cancel":
			case "expire":
				job.Status = "Closed"
				// TODO set transaction status on your databaase to 'failure'
				break
			case "pending":
				break
			}
			jobRepository.Update(job, tx)

			log.Printf("Order update processed successfully for order ID: %s", orderID)
			return nil
		})
		if err != nil {
			log.Println(err)
		}
		return nil
	})

	return &JobServiceImpl{
		jobRepository: jobRepository,
		//userRepository:        userRepository,
		//categoryRepository:    categoryRepository,
		dbConnection: dbConnection,
		//jobResourceRepository: jobResourceRepository,
		fileStorage: fileStorage,
		//userAddressRepository: userAddressRepository,
		validatorService: validatorService,
		//workerRepository:      workerRepository
		serviceDiscovery: serviceDiscovery,
		rabbitMQ:         rabbitMQ,
	}
}

func (jobService *JobServiceImpl) HandleCreate(ctx context.Context, userJwtClaims *dto.JwtClaimDto, createJobDto *jobDto.CreateJobDto, uploadedFiles []*multipart.FileHeader) *exception.ClientError {
	err := jobService.validatorService.ValidateStruct(createJobDto)
	jobService.validatorService.ParseValidationError(err)
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var jobModel model.Job
		var userAddress model.UserAddress
		userGrpcConnection, err := discovery.ServiceConnection(ctx, "userService", jobService.serviceDiscovery)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
		userAddressGrpcClient := userAddressGrpc.NewUserAddressServiceClient(userGrpcConnection)
		categoryServiceConnection, err := discovery.ServiceConnection(ctx, "categoryService", jobService.serviceDiscovery)
		categoryGrpcClient := category.NewCategoryServiceClient(categoryServiceConnection)
		userGrpcClient := user.NewUserServiceClient(userGrpcConnection)
		userModel, err := userGrpcClient.FindByIdentifier(ctx, &user.UserIdentifier{
			Email:       helper.SafeDereference(userJwtClaims.Email, ""),
			PhoneNumber: helper.SafeDereference(userJwtClaims.PhoneNumber, ""),
		})
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
		if createJobDto.AddressId == nil {
			queryResponse, err := userAddressGrpcClient.UserAddressStore(ctx, &userAddressGrpc.UserAddressCreateRequest{
				Latitude:  createJobDto.Latitude,
				Longitude: createJobDto.Longitude,
				UserId:    userModel.ID,
			})
			helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
			userAddress.ID = queryResponse.Id
		} else {
			queryResponse, err := userAddressGrpcClient.FindUserAddressById(ctx, &userAddressGrpc.UserAddressSearchRequest{
				UserId:        userModel.ID,
				UserAddressId: userAddress.ID,
			})
			helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
			userAddress.ID = queryResponse.Id
		}
		_, err = categoryGrpcClient.IsCategoryExists(ctx, &category.SearchCategoryRequest{CategoryId: createJobDto.CategoryId})
		if err != nil {
			exception.ThrowClientError(exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("category not found")))
		}
		mapper.MapJobDtoIntoJobModel(createJobDto, &jobModel)
		jobModel.UserId = userModel.ID
		jobModel.AddressId = userAddress.ID
		jobService.jobRepository.Store(&jobModel, gormTransaction)
		//uuidString := uuid.New().String()
		//var allFileName []string
		//for _, uploadedFile := range uploadedFiles {
		//	openedFile, _ := uploadedFile.Open()
		//	driverLicensePath := fmt.Sprintf("%s-%d-%s", uuidString, jobModel.ID, uploadedFile.Filename)
		//	_, err = jobService.fileStorage.UploadFile(openedFile, driverLicensePath)
		//	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("upload file failed")))
		//	allFileName = append(allFileName, uploadedFile.Filename)
		//}
		//resourceModel := mapper.MapStringIntoJobResourceModel(jobModel.ID, allFileName)
		//jobService.jobResourceRepository.BulkCreate(gormTransaction, resourceModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	return nil
}

func (jobService *JobServiceImpl) FindById(ctx context.Context, userEmail *string, jobId *uint64) *jobDto.JobResponseDto {
	var jobModel *model.Job
	var err error
	err = jobService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		userGrpcConnection, err := discovery.ServiceConnection(ctx, "userService", jobService.serviceDiscovery)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
		userAddressGrpcClient := user.NewUserServiceClient(userGrpcConnection)
		identifier, err := userAddressGrpcClient.FindByIdentifier(ctx, &user.UserIdentifier{
			Email:       helper.SafeDereference(userEmail, ""),
			PhoneNumber: "",
		})
		exception.ParseGrpcError(err)
		jobModel, err = jobService.jobRepository.FindVerifyById(gormTransaction, &identifier.ID, jobId)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return mapper.MapJobModelIntoJobResponseDto(jobModel)
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
