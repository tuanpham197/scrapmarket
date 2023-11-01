package service

import (
	"context"
	"sendo/internal/shop_category/service/entity"
	"sendo/internal/shop_category/service/request"
)

type ShopCategoryUseCase interface {
	CreateShopCategory(ctx context.Context, createShopCategory request.CreateShopCategory) (*request.CreateShopCategoryResponse, error)
	UpdateShopCategory(ctx context.Context, updateShopCategory request.UpdateShopCategory, shopId string, categoryId string) (*entity.ShopCategory, error)
	DeleteShopCategory(ctx context.Context, shopId string, categoryId string) (*entity.ShopCategory, error)
	GetListShopCategoryAndProduct(ctx context.Context, shopId string) (*[]entity.ShopCategory, error)
}

type ShopCategoryRepository interface {
	Insert(ctx context.Context, createShopCategory request.CreateShopCategory) (*entity.ShopCategory, error)
	UpdateShopCategory(ctx context.Context, updateShopCategory request.UpdateShopCategory, shopId string, categoryId string) (*entity.ShopCategory, error)
	DeleteShopCategory(ctx context.Context, shopId string, categoryId string) (*entity.ShopCategory, error)
	GetListShopCategoryAndProduct(ctx context.Context, shopId string) (*[]entity.ShopCategory, error)
}
