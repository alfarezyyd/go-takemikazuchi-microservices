package handler

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices-category/internal/category/service"
	grpcCategory "github.com/alfarezyyd/go-takemikazuchi-microservices-common/genproto/category"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryHandler struct {
	categoryService service.Service
	grpcCategory.UnimplementedCategoryServiceServer
}

func NewCategoryHandler(grpcServer *grpc.Server, categoryService service.Service) {
	categoryHandler := &CategoryHandler{
		categoryService: categoryService,
	}
	grpcCategory.RegisterCategoryServiceServer(grpcServer, categoryHandler)
}

func (categoryHandler *CategoryHandler) CreateCategory(ctx context.Context, categoryCreateRequest *grpcCategory.CreateCategoryRequest) (*grpcCategory.CommandCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (categoryHandler *CategoryHandler) UpdateCategory(ctx context.Context, categoryCreateRequest *grpcCategory.CreateCategoryRequest) (*grpcCategory.CommandCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
