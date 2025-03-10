package handler

import (
	"context"
	jobApplication "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/job_application"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/pkg/mapper"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/internal/service"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/job/pkg/dto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type JobApplicationHandler struct {
	jobApplicationService service.JobApplicationService
	jobApplication.UnimplementedJobApplicationServiceServer
}

func NewJobApplicationHandler(grpcServer *grpc.Server, jobApplicationService service.JobApplicationService,
) {
	jobApplicationHandler := &JobApplicationHandler{
		jobApplicationService: jobApplicationService,
	}
	jobApplication.RegisterJobApplicationServiceServer(grpcServer, jobApplicationHandler)
}

func (jobApplicationHandler *JobApplicationHandler) FindById(ctx context.Context, findIdByRequest *jobApplication.FindJobApplicationByIdRequest) (*jobApplication.JobApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindById not implemented")
}

func (jobApplicationHandler *JobApplicationHandler) FindAllApplication(ctx context.Context, findAllApplicationRequest *jobApplication.FindAllApplicationRequest) (*jobApplication.JobApplicationResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllApplication not implemented")
}
func (jobApplicationHandler *JobApplicationHandler) SelectApplication(ctx context.Context, selectApplicationRequest *jobApplication.SelectApplicationRequest) (*emptypb.Empty, error) {
	jobApplicationHandler.jobApplicationService.SelectApplication(ctx,
		mapper.MapUserJwtClaimGrpcIntoUserJwtClaim(selectApplicationRequest.UserJwtClaim),
		&dto.SelectApplicationDto{
			UserId: selectApplicationRequest.UserId,
			JobId:  selectApplicationRequest.JobId,
		})
	return nil, nil
}
func (jobApplicationHandler *JobApplicationHandler) HandleApply(ctx context.Context, jobApplicationApplyRequest *jobApplication.ApplyRequest) (*emptypb.Empty, error) {
	jobApplicationHandler.jobApplicationService.HandleApply(ctx, mapper.MapUserJwtClaimGrpcIntoUserJwtClaim(jobApplicationApplyRequest.UserJwtClaim), &dto.ApplyJobApplicationDto{
		JobId: jobApplicationApplyRequest.JobId,
	})
	return nil, nil
}
