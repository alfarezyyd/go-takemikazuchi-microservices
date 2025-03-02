package worker

import (
	"fmt"
	"github.com/google/uuid"
	"go-takemikazuchi-microservices/internal/model"
	"go-takemikazuchi-microservices/internal/storage"
	"go-takemikazuchi-microservices/internal/user"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
	validatorFeature "go-takemikazuchi-microservices/internal/validator"
	"go-takemikazuchi-microservices/internal/worker/dto"
	workerResource "go-takemikazuchi-microservices/internal/worker_resource"
	workerResourceDto "go-takemikazuchi-microservices/internal/worker_resource/dto"
	workerWallet "go-takemikazuchi-microservices/internal/worker_wallet"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"go-takemikazuchi-microservices/pkg/mapper"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	workerRepository         Repository
	userRepository           user.Repository
	validatorService         validatorFeature.Service
	dbConnection             *gorm.DB
	workerWalletRepository   workerWallet.Repository
	workerResourceRepository workerResource.Repository
	fileStorage              storage.FileStorage
}

func NewService(
	workerRepository Repository,
	validatorService validatorFeature.Service,
	dbConnection *gorm.DB,
	userRepository user.Repository,
	workerWalletRepository workerWallet.Repository,
	workerResourceRepository workerResource.Repository,
	fileStorage storage.FileStorage,
) *ServiceImpl {
	return &ServiceImpl{
		workerRepository:         workerRepository,
		validatorService:         validatorService,
		dbConnection:             dbConnection,
		workerWalletRepository:   workerWalletRepository,
		workerResourceRepository: workerResourceRepository,
		fileStorage:              fileStorage,
		userRepository:           userRepository,
	}
}

func (workerService *ServiceImpl) Create(userJwtClaim *userDto.JwtClaimDto, createWorkerDto *dto.CreateWorkerDto, createWorkerWalletDocumentDto *workerResourceDto.CreateWorkerWalletDocumentDto) {
	err := workerService.validatorService.ValidateStruct(createWorkerDto)
	workerService.validatorService.ParseValidationError(err)
	err = workerService.validatorService.ValidateStruct(createWorkerWalletDocumentDto)
	workerService.validatorService.ParseValidationError(err)
	err = workerService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var workerModel model.Worker
		var workerWalletModel model.WorkerWallet
		var userModel model.User
		workerService.userRepository.FindUserByEmail(userJwtClaim.Email, &userModel, gormTransaction)
		mapper.MapCreateWorkerDtoIntoWorkerModel(&workerModel, createWorkerDto)
		mapper.MapCreateWorkerWalletDtoIntoWorkerWalletModel(&workerWalletModel, &createWorkerDto.WalletInformation)
		workerModel.UserId = userModel.ID
		workerService.workerRepository.Store(gormTransaction, &workerModel)
		workerWalletModel.WorkerID = workerModel.ID
		workerService.workerWalletRepository.Store(gormTransaction, &workerWalletModel)
		driverLicenseFile, _ := createWorkerWalletDocumentDto.DriverLicense.Open()
		identityCardFile, _ := createWorkerWalletDocumentDto.IdentityCard.Open()
		policeCertificateFile, _ := createWorkerWalletDocumentDto.PoliceCertificate.Open()
		uuidString := uuid.New().String()
		driverLicensePath := fmt.Sprintf("%s-%s-%s", uuidString, "driverLicense", createWorkerWalletDocumentDto.DriverLicense.Filename)
		policeCertificatePath := fmt.Sprintf("%s-%s-%s", uuidString, "policeCertificate", createWorkerWalletDocumentDto.PoliceCertificate.Filename)
		identityCardPath := fmt.Sprintf("%s-%s-%s", uuidString, "identityCard", createWorkerWalletDocumentDto.IdentityCard.Filename)
		workerService.fileStorage.UploadFile(driverLicenseFile, driverLicensePath)
		workerService.fileStorage.UploadFile(policeCertificateFile, policeCertificatePath)
		workerService.fileStorage.UploadFile(identityCardFile, identityCardPath)
		workerResourcesModel := mapper.MapStringIntoWorkerResourceModel(workerModel.ID,
			[]string{driverLicensePath, policeCertificatePath, identityCardPath},
			[]string{"Driver License", "Police Certificate", "Identity Card"},
			3,
		)
		workerService.workerResourceRepository.BulkStore(gormTransaction, workerResourcesModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
