package handler

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/withdrawal"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WithdrawalHandler struct {
	withdrawalService service.WithdrawalService
	withdrawal.UnimplementedWithdrawalServiceServer
}

func NewWithdrawalHandler(grpcServer *grpc.Server, withdrawalService service.WithdrawalService) {
	withdrawalHandler := &WithdrawalHandler{
		withdrawalService: withdrawalService,
	}
	withdrawal.RegisterWithdrawalServiceServer(grpcServer, withdrawalHandler)
}

func (withdrawalHandler *WithdrawalHandler) Create(ctx context.Context, createWithdrawalDto *withdrawal.CreateWithdrawal) (*emptypb.Empty, error) {
	withdrawalHandler.withdrawalService.Create(ctx,
		mapper.MapUserJwtClaimGrpcIntoUserJwtClaim(createWithdrawalDto.UserJwtClaim),
		&dto.CreateWithdrawalDto{
			WalletId: createWithdrawalDto.WalletId,
			Amount:   createWithdrawalDto.Amount,
		})
	return &emptypb.Empty{}, nil
}
