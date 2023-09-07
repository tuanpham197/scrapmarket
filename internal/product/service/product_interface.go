package service

import (
	"context"
	"sendo/internal/product/service/entity"
	"sendo/internal/product/service/request"
	"sendo/pkg/utils/paginations"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, createProduct request.CreateProductRaw, shopId string) (*request.CreateProductResponse, error)
	GetList(ctx context.Context, filter *request.FilterRequest) (*paginations.Pagination, error)
}

type ProductRepository interface {
	InsertProduct(ctx context.Context, createProduct request.CreateProductRaw, shopId string) (*entity.Product, error)
	GetList(ctx context.Context, filter *request.FilterRequest) (*paginations.Pagination, error)
}
