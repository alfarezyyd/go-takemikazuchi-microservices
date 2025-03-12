package service

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/internal/repository"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/pkg/dto"
	"gorm.io/gorm"
)

type WorkerWalletServiceImpl struct {
	validatorService       validatorFeature.Service
	dbConnection           *gorm.DB
	workerWalletRepository repository.WorkerWalletRepository
	serviceRegistry        discovery.ServiceRegistry
}

func NewWorkerWalletServiceImpl(
	validatorService validatorFeature.Service,
	dbConnection *gorm.DB,
	workerWalletRepository repository.WorkerWalletRepository,
	serviceRegistry discovery.ServiceRegistry,
) *WorkerWalletServiceImpl {
	return &WorkerWalletServiceImpl{
		validatorService:       validatorService,
		dbConnection:           dbConnection,
		workerWalletRepository: workerWalletRepository,
		//workerResourceRepository: workerResourceRepository,
		//userRepository:           userRepository,
		serviceRegistry: serviceRegistry,
	}
}

func (workerWalletServiceImpl *WorkerWalletServiceImpl) Create(ctx context.Context, userJwtClaim *userDto.JwtClaimDto, createWorkerDto *dto.CreateWorkerDto) {
	return
}

func (workerWalletServiceImpl *WorkerWalletServiceImpl) FindById(ctx context.Context, walletId *uint64) *dto.ResponseWorkerWalletDto {
	var workerWallet *model.WorkerWallet
	err := workerWalletServiceImpl.dbConnection.Transaction(func(tx *gorm.DB) error {
		workerWallet = workerWalletServiceImpl.workerWalletRepository.FindById(tx, walletId)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return mapper.MapWorkerWalletModelIntoWorkerWalletResponse(workerWallet)

}
