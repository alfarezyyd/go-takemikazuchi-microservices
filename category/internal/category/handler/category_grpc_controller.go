package handler

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/category/internal/category/service"
	categoryDto "github.com/alfarezyyd/go-takemikazuchi-microservices/category/pkg/dto"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/common/discovery"
	grpcCategory "github.com/alfarezyyd/go-takemikazuchi-microservices/common/genproto/category"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryHandler struct {
	categoryService service.CategoryService
	serviceRegistry discovery.ServiceRegistry
	grpcCategory.UnimplementedCategoryServiceServer
}

func NewCategoryHandler(grpcServer *grpc.Server, categoryService service.CategoryService) {
	categoryHandler := &CategoryHandler{
		categoryService: categoryService,
	}
	grpcCategory.RegisterCategoryServiceServer(grpcServer, categoryHandler)
}

func (categoryHandler *CategoryHandler) FindAll(ctx context.Context, emptyProtobuf *emptypb.Empty) (*grpcCategory.QueryCategoryResponses, error) {
	allCategory := categoryHandler.categoryService.FindAll()
	return allCategory, nil
}
func (categoryHandler *CategoryHandler) HandleCreate(ctx context.Context, createCategoryRequest *grpcCategory.CreateCategoryRequest) (*grpcCategory.CommandCategoryResponse, error) {
	categoryHandler.categoryService.HandleCreate(ctx, &dto.JwtClaimDto{
		Email:       createCategoryRequest.UserJwtClaim.Email,
		PhoneNumber: createCategoryRequest.UserJwtClaim.PhoneNumber,
	}, &categoryDto.CreateCategoryDto{
		Name:        createCategoryRequest.Name,
		Description: createCategoryRequest.Description,
	})
	return &grpcCategory.CommandCategoryResponse{
		IsSuccess: true,
	}, nil
}
func (categoryHandler *CategoryHandler) HandleUpdate(ctx context.Context, updateCategoryRequest *grpcCategory.UpdateCategoryRequest) (*grpcCategory.CommandCategoryResponse, error) {
	return nil, nil
}
func (categoryHandler *CategoryHandler) HandleDelete(ctx context.Context, deleteCategoryRequest *grpcCategory.DeleteCategoryRequest) (*grpcCategory.CommandCategoryResponse, error) {
	return nil, nil
}

func (categoryHandler *CategoryHandler) IsCategoryExists(ctx context.Context, searchCategoryRequest *grpcCategory.SearchCategoryRequest) (*grpcCategory.CategoryExistsResponse, error) {
	isCategoryExists, _ := categoryHandler.categoryService.IsCategoryExists(searchCategoryRequest.CategoryId)
	return &grpcCategory.CategoryExistsResponse{
		IsExists: isCategoryExists,
	}, nil
}
