package service

import (
	"context"
	"fmt"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/worker"
	workerWallet "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/worker_wallet"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/model"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	validatorFeature "github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/validator"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/repository"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"gorm.io/gorm"
	"net/http"
)

type WithdrawalServiceImpl struct {
	validatorService     validatorFeature.Service
	dbConnection         *gorm.DB
	withdrawalRepository repository.WithdrawalRepository
	userRepository       repository.UserRepository
	serviceRegistry      discovery.ServiceRegistry
	//workerRepository     worker.Repository
	//walletRepository     workerWalletFeature.Repository
}

func NewWithdrawalService(
	validatorService validatorFeature.Service,
	withdrawalRepository repository.WithdrawalRepository,
	dbConnection *gorm.DB,
	userRepository repository.UserRepository,
	serviceRegistry discovery.ServiceRegistry,
	// workerRepository worker.Repository,
	// walletRepository workerWalletFeature.Repository
) *WithdrawalServiceImpl {
	return &WithdrawalServiceImpl{
		validatorService:     validatorService,
		withdrawalRepository: withdrawalRepository,
		dbConnection:         dbConnection,
		userRepository:       userRepository,
		serviceRegistry:      serviceRegistry,
		//workerRepository:     workerRepository,
		//walletRepository:     walletRepository,
	}
}

func (withdrawalService *WithdrawalServiceImpl) FindAll(userJwtClaims *userDto.JwtClaimDto) []model.Withdrawal {
	//		var withdrawalsModel []model.Withdrawal
	//		err := withdrawalService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
	//			var userModel model.User
	//			withdrawalService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
	//			if userModel.Role != "Admin" {
	//				exception.ThrowClientError(exception.NewClientError(http.StatusUnauthorized, exception.ErrUnauthorized, errors.New("only admin can do the ops")))
	//			}
	//			withdrawalsModel = withdrawalService.withdrawalRepository.FindAll(gormTransaction)
	//			return nil
	//		})
	//		helper.CheckErrorOperation(err, exception.ParseGormError(err))
	//		return withdrawalsModel
	return nil
}
func (withdrawalService *WithdrawalServiceImpl) Create(ctx context.Context, userJwtClaims *userDto.JwtClaimDto, createWithdrawalDto *dto.CreateWithdrawalDto) {
	err := withdrawalService.validatorService.ValidateStruct(userJwtClaims)
	withdrawalService.validatorService.ParseValidationError(err)
	fmt.Println("CHECKPOINT 1")
	err = withdrawalService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
		var userModel model.User
		withdrawalService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
		workerServiceConnection, err := discovery.ServiceConnection(ctx, "workerService", withdrawalService.serviceRegistry)
		workerGrpcClient := worker.NewWorkerServiceClient(workerServiceConnection)
		workerModel, err := workerGrpcClient.FindById(ctx, &worker.SearchWorkerRequest{
			UserId: userModel.ID,
		})
		fmt.Println(err)
		helper.CheckErrorOperation(err, exception.ParseGormError(err))
		workerWalletGrpcClient := workerWallet.NewWorkerWalletServiceClient(workerServiceConnection)
		fmt.Println(workerModel)
		workerWalletResponse, err := workerWalletGrpcClient.FindById(ctx, &workerWallet.SearchRequest{
			UserId: workerModel.ID,
		})
		fmt.Println(err)
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
		withdrawalModel := mapper.MapWithdrawalDtoIntoWithdrawalModel(createWithdrawalDto)
		withdrawalModel.WorkerId = workerModel.ID
		withdrawalModel.WalletId = workerWalletResponse.ID
		withdrawalService.withdrawalRepository.Create(gormTransaction, withdrawalModel)
		return nil
	})
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}

func (withdrawalService *WithdrawalServiceImpl) Update(userJwtClaims *userDto.JwtClaimDto, withdrawalId *string) {
	//	err := withdrawalService.validatorService.ValidateVar(withdrawalId, "required|gt=0")
	//	withdrawalService.validatorService.ParseValidationError(err)
	//	err = withdrawalService.dbConnection.Transaction(func(gormTransaction *gorm.DB) error {
	//		var userModel model.User
	//		parsedWithdrawalId, err := strconv.ParseUint(*withdrawalId, 10, 64)
	//		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("withdrawal id not valid")))
	//		withdrawalService.userRepository.FindUserByEmail(userJwtClaims.Email, &userModel, gormTransaction)
	//		withdrawalModel, err := withdrawalService.withdrawalRepository.FindById(gormTransaction, &parsedWithdrawalId)
	//		helper.CheckErrorOperation(err, exception.ParseGormError(err))
	//		withdrawalService.withdrawalRepository.Update(gormTransaction, withdrawalModel)
	//		return nil
	//	})
	//	helper.CheckErrorOperation(err, exception.ParseGormError(err))
}
