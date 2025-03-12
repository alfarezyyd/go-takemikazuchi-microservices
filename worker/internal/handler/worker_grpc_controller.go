package handler

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	grpcWorker "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/worker"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/worker/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WorkerHandler struct {
	workerService   service.WorkerService
	serviceRegistry discovery.ServiceRegistry
	grpcWorker.UnimplementedWorkerServiceServer
}

func NewWorkerHandler(grpcServer *grpc.Server, workerService service.WorkerService) {
	workerHandler := &WorkerHandler{
		workerService: workerService,
	}
	grpcWorker.RegisterWorkerServiceServer(grpcServer, workerHandler)
}

func (workerHandler *WorkerHandler) Create(ctx context.Context, createWorkerRequest *grpcWorker.CreateWorkerRequest) (*emptypb.Empty, error) {
	userJwtClaimDto := mapper.MapUserJwtClaimGrpcIntoUserJwtClaim(createWorkerRequest.UserJwtClaim)
	createWorkerDto := mapper.MapCreateWorkerRequestIntoCreateWorkerDto(createWorkerRequest)
	workerHandler.workerService.Create(ctx, userJwtClaimDto, createWorkerDto)
	return nil, nil
}

func (workerHandler *WorkerHandler) FindById(ctx context.Context, searchWorkRequest *grpcWorker.SearchWorkerRequest) (*grpcWorker.WorkerResponse, error) {
	workerResponseDto := workerHandler.workerService.FindById(ctx, &searchWorkRequest.UserId)
	return mapper.MapWorkerResponseDtoIntoWorkerGrpc(workerResponseDto), nil
}
