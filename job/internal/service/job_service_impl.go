package service

import (
	"context"
	"errors"
	"fmt"
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
		//userAddressRepository: userAddressRepository,
		validatorService: validatorService,
		//workerRepository:      workerRepository
		serviceDiscovery: serviceDiscovery,
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
		fmt.Println("CHECKPOINT 3")
		isCategoryExists, err := categoryGrpcClient.IsCategoryExists(ctx, &category.SearchCategoryRequest{CategoryId: createJobDto.CategoryId})
		if err != nil {
			exception.ThrowClientError(exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("category not found")))
		}
		fmt.Println(isCategoryExists)
		mapper.MapJobDtoIntoJobModel(createJobDto, &jobModel)
		jobModel.UserId = userModel.ID
		jobModel.AddressId = userAddress.ID
		jobService.jobRepository.Store(&jobModel, gormTransaction)
		fmt.Println("CHECKPOINT 4")
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
	fmt.Println(err)
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
