package withdrawal

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-takemikazuchi-api/internal/model"
	"go-takemikazuchi-api/internal/user"
	userDto "go-takemikazuchi-api/internal/user/dto"
	"go-takemikazuchi-api/internal/withdrawal/dto"
	"go-takemikazuchi-api/internal/worker"
	workerWalletFeature "go-takemikazuchi-api/internal/worker_wallet"
	"go-takemikazuchi-api/pkg/exception"
	"go-takemikazuchi-api/pkg/helper"
	"go-takemikazuchi-api/pkg/mapper"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	validationService    *validator.Validate
	engTranslator        ut.Translator
	dbConnection         *gorm.DB
	withdrawalRepository Repository
	userRepository       user.Repository
	workerRepository     worker.Repository
	walletRepository     workerWalletFeature.Repository
}

func NewServiceImpl(
	validationService *validator.Validate,
	engTranslator ut.Translator,
	withdrawalRepository Repository) *ServiceImpl {
	return &ServiceImpl{
		validationService:    validationService,
		engTranslator:        engTranslator,
		withdrawalRepository: withdrawalRepository,
	}
}

func (withdrawalService *ServiceImpl) Create(userJwtClaims *userDto.JwtClaimDto, createWithdrawalDto *dto.CreateWithdrawalDto) {
	err := withdrawalService.validationService.Struct(userJwtClaims)
	exception.ParseValidationError(err, withdrawalService.engTranslator)
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
