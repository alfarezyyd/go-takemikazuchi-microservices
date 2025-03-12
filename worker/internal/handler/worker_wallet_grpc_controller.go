package handler

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	workerWallet "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/worker_wallet"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/internal/service"
	"google.golang.org/grpc"
)

type WorkerWalletHandler struct {
	workerWalletService service.WorkerWalletService
	serviceRegistry     discovery.ServiceRegistry
	workerWallet.UnimplementedWorkerWalletServiceServer
}

func NewWorkerWalletHandler(grpcServer *grpc.Server, workerWalletService service.WorkerWalletService) {
	workerHandler := &WorkerWalletHandler{
		workerWalletService: workerWalletService,
	}
	workerWallet.RegisterWorkerWalletServiceServer(grpcServer, workerHandler)
}

func (workerWalletHandler *WorkerWalletHandler) FindById(ctx context.Context, searchRequest *workerWallet.SearchRequest) (*workerWallet.WorkerWalletResponse, error) {
	workerWalletResponseDto := workerWalletHandler.workerWalletService.FindById(ctx, &searchRequest.UserId)
	workerWalletGrpcResponse := mapper.MapWorkerWalletResponseIntoWorkerWalletGrpcResponse(workerWalletResponseDto)
	return workerWalletGrpcResponse, nil
}
