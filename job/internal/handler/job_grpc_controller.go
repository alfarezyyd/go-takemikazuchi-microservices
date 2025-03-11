package handler

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job"
	grpcJob "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type JobHandler struct {
	jobService service.JobService
	job.UnimplementedJobServiceServer
}

func NewJobHandler(grpcServer *grpc.Server, jobService service.JobService,
) *JobHandler {
	jobHandler := &JobHandler{
		jobService: jobService,
	}
	grpcJob.RegisterJobServiceServer(grpcServer, jobHandler)
	return jobHandler
}

func (jobHandler *JobHandler) FindAll(ctx context.Context, emptyProto *emptypb.Empty) (*job.JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (jobHandler *JobHandler) HandleCreate(ctx context.Context, createJobRequest *job.CreateJobRequest) (*emptypb.Empty, error) {
	createJobGrpc := mapper.MapCreateJobGrpcIntoCreateJobDto(createJobRequest)
	userJwtClaimGrpc := mapper.MapUserJwtClaimGrpcIntoUserJwtClaim(createJobRequest.UserJwtClaim)
	jobHandler.jobService.HandleCreate(ctx, userJwtClaimGrpc, createJobGrpc, nil)
	return nil, nil
}
func (jobHandler *JobHandler) HandleUpdate(ctx context.Context, updateJobRequest *job.UpdateJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleUpdate not implemented")
}
func (jobHandler *JobHandler) HandleDelete(ctx context.Context, deleteJobRequest *job.DeleteJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleDelete not implemented")
}
func (jobHandler *JobHandler) HandleRequestCompleted(ctx context.Context, jobCompleteRequest *job.JobCompleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleRequestCompleted not implemented")
}

func (jobHandler *JobHandler) FindById(ctx context.Context, findByIdRequest *job.FindByIdRequest) (*job.JobModel, error) {
	jobModel := jobHandler.jobService.FindById(ctx, &findByIdRequest.UserEmail, &findByIdRequest.JobId)
	return mapper.MapJobResponseIntoJobModel(jobModel), nil
}
