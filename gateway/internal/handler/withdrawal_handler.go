package handler

import (
	"context"
	"errors"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/exception"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/withdrawal"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/helper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	userDto "github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type WithdrawalHandler struct {
	serviceDiscovery discovery.ServiceRegistry
}

func NewWithdrawalHandler(serviceDiscovery discovery.ServiceRegistry) *WithdrawalHandler {
	return &WithdrawalHandler{
		serviceDiscovery: serviceDiscovery,
	}

}

//func (withdrawalHandler *WithdrawalHandler) FindAll(ginContext *gin.Context) {
//	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
//	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
//	defer cancelFunc()
//	grpcConnection, _ := discovery.ServiceConnection(timeoutCtx, "userService", withdrawalHandler.serviceDiscovery)
//	userClient := job.NewJobServiceClient(grpcConnection)
//
//	withdrawalsModel, _ := userClient.FindAll(timeoutCtx, &emptypb.Empty{})
//	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Success", withdrawalsModel))
//}

func (withdrawalHandler *WithdrawalHandler) Create(ginContext *gin.Context) {
	var createWithdrawalDto dto.CreateWithdrawalDto
	err := ginContext.ShouldBindBodyWithJSON(&createWithdrawalDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrBadRequest, errors.New("error parsing body")))
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	defer cancelFunc()
	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
	grpcConnection, err := discovery.ServiceConnection(timeoutCtx, "userService", withdrawalHandler.serviceDiscovery)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusInternalServerError, exception.ErrInternalServerError, err))
	withdrawalClient := withdrawal.NewWithdrawalServiceClient(grpcConnection)

	_, err = withdrawalClient.Create(timeoutCtx, &withdrawal.CreateWithdrawal{
		UserJwtClaim: mapper.MapUserJwtClaimIntoUserJwtClaimGrpc(userJwtClaim),
		WalletId:     createWithdrawalDto.WalletId,
		Amount:       createWithdrawalDto.Amount,
	})
	exception.ParseGrpcError(err)
	ginContext.JSON(http.StatusCreated, helper.WriteSuccess("Success", nil))
}

//func (withdrawalHandler *WithdrawalHandler) Update(ginContext *gin.Context) {
//	withdrawalId := ginContext.Param("withdrawalId")
//	userJwtClaim := ginContext.MustGet("claims").(*userDto.JwtClaimDto)
//	withdrawalHandler.withdrawalService.Update(userJwtClaim, &withdrawalId)
//	ginContext.JSON(http.StatusOK, helper.WriteSuccess("Success", nil))
//}
