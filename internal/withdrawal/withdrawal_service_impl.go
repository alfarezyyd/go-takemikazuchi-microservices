package withdrawal

import (
	"errors"
	"go-takemikazuchi-microservices/internal/model"
	"go-takemikazuchi-microservices/internal/user"
	userDto "go-takemikazuchi-microservices/internal/user/dto"
	validatorFeature "go-takemikazuchi-microservices/internal/validator"
	"go-takemikazuchi-microservices/internal/withdrawal/dto"
	"go-takemikazuchi-microservices/internal/worker"
	workerWalletFeature "go-takemikazuchi-microservices/internal/worker_wallet"
	"go-takemikazuchi-microservices/pkg/exception"
	"go-takemikazuchi-microservices/pkg/helper"
	"go-takemikazuchi-microservices/pkg/mapper"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ServiceImpl struct {
	validatorService     validatorFeature.Service
	dbConnection         *gorm.DB
	withdrawalRepository Repository
	userRepository       user.Repository
	workerRepository     worker.Repository
	walletRepository     workerWalletFeature.Repository
}

func NewService(
	validatorService validatorFeature.Service,
	withdrawalRepository Repository,
	dbConnection *gorm.DB,
	userRepository user.Repository,
	workerRepository worker.Repository,
	walletRepository workerWalletFeature.Repository) *ServiceImpl {
	return &ServiceImpl{
		validatorService:     validatorService,
		withdrawalRepository: withdrawalRepository,
		dbConnection:         dbConnection,
		userRepository:       userRepository,
		workerRepository:     workerRepository,
		walletRepository:     walletRepository,
	}
}
func (withdrawalService *ServiceImpl) FindAll(userJwtClaims *userDto.JwtClaimDto) []model.Withdrawal {
	var withdrawalsModel []model.Withdrawal
	err := withdrawalService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		withdrawalService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		if userModel.Role != "Admin" {
			exception.ThrowClientError(exception.NewClientError(http.StatusUnauthorized, exception.ErrUnauthorized, errors.New("only admin can do the ops")))
		}
		withdrawalsModel = withdrawalService.withdrawalRepository.FindAll(gormTransaction)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	return withdrawalsModel
}

func (withdrawalService *ServiceImpl) Create(userJwtClaims *userDto.JwtClaimDto, createWithdrawalDto *dto.CreateWithdrawalDto) {
	err := withdrawalService.validatorService.ValidateStruct(userJwtClaims)
	withdrawalService.validatorService.ParseValidationError(err)
	err = withdrawalService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var withdrawalModel model.Withdrawal
		var userModel model.User
		withdrawalService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		workerModel, err := withdrawalService.workerRepository.FindById(gormTransaction, &userModel.ID)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		withdrawalService.walletRepository.FindById(gormTransaction, &createWithdrawalDto.WalletId)
		mapper.MapCreateWithdrawalDtoIntoWithdrawalModel(createWithdrawalDto, &withdrawalModel)
		withdrawalModel.WorkerId = workerModel.ID
		withdrawalService.withdrawalRepository.Create(gormTransaction, &withdrawalModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (withdrawalService *ServiceImpl) Update(userJwtClaims *userDto.JwtClaimDto, withdrawalId *string) {
	err := withdrawalService.validatorService.ValidateVar(withdrawalId, "required|gt=0")
	withdrawalService.validatorService.ParseValidationError(err)
	err = withdrawalService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		parsedWithdrawalId, err := strconv.ParseUint(*withdrawalId, 10, 64)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("withdrawal id not valid")))
		withdrawalService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		withdrawalModel, err := withdrawalService.withdrawalRepository.FindById(gormTransaction, &parsedWithdrawalId)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		withdrawalService.withdrawalRepository.Update(gormTransaction, withdrawalModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
