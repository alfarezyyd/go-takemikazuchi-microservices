package handler

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/transaction"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/internal/service"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/payment/pkg/dto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TransactionHandler struct {
	transactionService service.TransactionService
	serviceRegistry    discovery.ServiceRegistry
	transaction.UnimplementedTransactionServiceServer
}

func NewTransactionHandler(grpcServer *grpc.Server, transactionService service.TransactionService) {
	transactionHandler := &TransactionHandler{
		transactionService: transactionService,
	}
	transaction.RegisterTransactionServiceServer(grpcServer, transactionHandler)
}

func (transactionHandler *TransactionHandler) Create(ctx context.Context, createTransactionRequest *transaction.CreateTransactionRequest) (*transaction.TransactionResponse, error) {
	snapToken := transactionHandler.transactionService.Create(ctx, mapper.MapUserJwtClaimGrpcIntoUserJwtClaim(createTransactionRequest.UserJwtClaim), &dto.CreateTransactionDto{
		JobId:       createTransactionRequest.JobId,
		ApplicantId: createTransactionRequest.ApplicantId,
	})
	return &transaction.TransactionResponse{SnapToken: snapToken}, nil
}
func (transactionHandler *TransactionHandler) PostPayment(ctx context.Context, postPaymentRequest *transaction.PostPaymentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostPayment not implemented")
}
