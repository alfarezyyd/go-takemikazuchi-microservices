package service

import (
	"context"
	"github.com/alfarezyyd/go-takemikazuchi-microservices/user/pkg/dto"
)

type UserAddressService interface {
	Create(ctx context.Context, createUserAddressDto *dto.CreateUserAddressDto) uint64
	FindById(ctx context.Context, searchUserAddressDto *dto.SearchUserAddressDto)
}
