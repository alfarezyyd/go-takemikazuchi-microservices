package worker

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/internal/storage"
	"go-takemikazuchi-api/internal/worker/dto"
	workerResource "go-takemikazuchi-api/internal/worker_resource"
	workerResourceDto "go-takemikazuchi-api/internal/worker_resource/dto"
	workerWallet "go-takemikazuchi-api/internal/worker_wallet"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"go-takemikazuchi-api/pkg/mapper"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	workerRepository         Repository
	validatorInstance        *validator.Validate
	engTranslator            ut.Translator
	dbConnection             *gorm.DB
	workerWalletRepository   workerWallet.Repository
	workerResourceRepository workerResource.Repository
	fileStorage              storage.FileStorage
}

func NewService(
	workerRepository Repository,
	validatorInstance *validator.Validate,
	engTranslator ut.Translator,
	dbConnection *gorm.DB,
	workerWalletRepository workerWallet.Repository,
	workerResourceRepository workerResource.Repository,
	fileStorage storage.FileStorage,
) *ServiceImpl {
	return &ServiceImpl{
		workerRepository:         workerRepository,
		validatorInstance:        validatorInstance,
		engTranslator:            engTranslator,
		dbConnection:             dbConnection,
		workerWalletRepository:   workerWalletRepository,
		workerResourceRepository: workerResourceRepository,
		fileStorage:              fileStorage,
	}
}

func (workerService *ServiceImpl) Create(createWorkerDto *dto.CreateWorkerDto, createWorkerWalletDocumentDto *workerResourceDto.CreateWorkerWalletDocumentDto) {
	err := workerService.validatorInstance.Struct(createWorkerDto)
	exception.ParseValidationError(err, workerService.engTranslator)
	err = workerService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var workerModel model.Worker
		var workerWalletModel model.WorkerWallet
		mapper.MapCreateWorkerDtoIntoWorkerModel(&workerModel, createWorkerDto)
		mapper.MapCreateWorkerWalletDtoIntoWorkerWalletModel(&workerWalletModel, &createWorkerDto.WalletInformation)
		workerService.workerRepository.Store(gormTransaction, &workerModel)
		workerService.workerWalletRepository.Store(gormTransaction, &workerWalletModel)
		driverLicenseFile, _ := createWorkerWalletDocumentDto.DriverLicense.Open()
		identityCardFile, _ := createWorkerWalletDocumentDto.IdentityCard.Open()
		policeCertificateFile, _ := createWorkerWalletDocumentDto.PoliceCertificate.Open()
		defer driverLicenseFile.Close()
		defer identityCardFile.Close()
		defer policeCertificateFile.Close()
		uuidString := uuid.New().String()
		driveLicensePath, _ := workerService.fileStorage.UploadFile(driverLicenseFile, fmt.Sprintf("%s-%s", uuidString, "driverLicense"))
		policeCertificatePath, _ := workerService.fileStorage.UploadFile(policeCertificateFile, fmt.Sprintf("%s-%s", uuidString, "driverLicense"))
		identityCardPath, _ := workerService.fileStorage.UploadFile(identityCardFile, fmt.Sprintf("%s-%s", uuidString, "driverLicense"))
		workerResourcesModel := mapper.MapStringIntoWorkerResourceModel(workerModel.ID,
			[]string{driveLicensePath, policeCertificatePath, identityCardPath},
			[]string{"Identity Card", "Police Certificate", "Driver License"},
			3,
		)
		workerService.workerResourceRepository.BulkStore(gormTransaction, workerResourcesModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
