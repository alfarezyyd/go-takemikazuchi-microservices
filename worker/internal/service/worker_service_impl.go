package service

import (
	"context"
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/user"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/storage"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/internal/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/pkg/dto"
	"gorm.io/gorm"
	"net/http"
)

type WorkerServiceImpl struct {
	workerRepository repository.WorkerRepository
	//userRepository           user.Repository
	validatorService       validatorFeature.Service
	dbConnection           *gorm.DB
	workerWalletRepository repository.WorkerWalletRepository
	//workerResourceRepository workerResource.Repository
	fileStorage     storage.FileStorage
	serviceRegistry discovery.ServiceRegistry
}

func NewWorkerService(
	workerRepository repository.WorkerRepository,
	validatorService validatorFeature.Service,
	dbConnection *gorm.DB,
	//userRepository user.Repository,
	//workerWalletRepository workerWallet.Repository,
	//workerResourceRepository workerResource.Repository,
	fileStorage storage.FileStorage,
	serviceRegistry discovery.ServiceRegistry,

) *WorkerServiceImpl {
	return &WorkerServiceImpl{
		workerRepository: workerRepository,
		validatorService: validatorService,
		dbConnection:     dbConnection,
		//workerWalletRepository:   workerWalletRepository,
		//workerResourceRepository: workerResourceRepository,
		fileStorage: fileStorage,
		//userRepository:           userRepository,
		serviceRegistry: serviceRegistry,
	}
}

func (workerService *WorkerServiceImpl) Create(ctx context.Context, userJwtClaim *userDto.JwtClaimDto, createWorkerDto *dto.CreateWorkerDto) {
	err := workerService.validatorService.ValidateStruct(createWorkerDto)
	workerService.validatorService.ParseValidationError(err)
	workerService.validatorService.ParseValidationError(err)
	userGrpcConnection, err := discovery.ServiceConnection(ctx, "userService", workerService.serviceRegistry)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, errors.New("user service down")))
	userGrpcClient := user.NewUserServiceClient(userGrpcConnection)
	err = workerService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var workerModel model.Worker
		var workerWalletModel model.WorkerWallet
		userModel, err := userGrpcClient.FindByIdentifier(ctx, &user.UserIdentifier{
			Email:       userJwtClaim.Email,
			PhoneNumber: userJwtClaim.PhoneNumber,
		})
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusNotFound, exception.ErrNotFound, errors.New("user not found")))
		mapper.MapCreateWorkerDtoIntoWorkerModel(&workerModel, createWorkerDto)
		mapper.MapCreateWorkerWalletDtoIntoWorkerWalletModel(&workerWalletModel, &createWorkerDto.WalletInformation)
		workerModel.UserId = userModel.ID
		workerService.workerRepository.Store(gormTransaction, &workerModel)
		workerWalletModel.WorkerID = workerModel.ID
		workerService.workerWalletRepository.Store(gormTransaction, &workerWalletModel)
		//driverLicenseFile, _ := createWorkerWalletDocumentDto.DriverLicense.Open()
		//identityCardFile, _ := createWorkerWalletDocumentDto.IdentityCard.Open()
		//policeCertificateFile, _ := createWorkerWalletDocumentDto.PoliceCertificate.Open()
		//uuidString := uuid.New().String()
		////driverLicensePath := fmt.Sprintf("%s-%s-%s", uuidString, "driverLicense", createWorkerWalletDocumentDto.DriverLicense.Filename)
		//policeCertificatePath := fmt.Sprintf("%s-%s-%s", uuidString, "policeCertificate", createWorkerWalletDocumentDto.PoliceCertificate.Filename)
		//identityCardPath := fmt.Sprintf("%s-%s-%s", uuidString, "identityCard", createWorkerWalletDocumentDto.IdentityCard.Filename)
		//workerService.fileStorage.UploadFile(driverLicenseFile, driverLicensePath)
		//workerService.fileStorage.UploadFile(policeCertificateFile, policeCertificatePath)
		//workerService.fileStorage.UploadFile(identityCardFile, identityCardPath)
		//workerResourcesModel := mapper.MapStringIntoWorkerResourceModel(workerModel.ID,
		//	[]string{driverLicensePath, policeCertificatePath, identityCardPath},
		//	[]string{"Driver License", "Police Certificate", "Identity Card"},
		//	3,
		//)
		//workerService.workerResourceRepository.BulkStore(gormTransaction, workerResourcesModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
