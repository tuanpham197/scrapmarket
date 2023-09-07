package service

import (
	"context"
	"sendo/internal/shop/service/entity"
	"sendo/internal/shop/service/request"
)

type ShopUseCase interface {
	ShopRegister(ctx context.Context, shopRegister request.ShopRegister) (*request.ShopRegisterResponse, error)
	GetShopInfo(ctx context.Context, userId string) (*entity.Shop, error)
}

type ShopRepository interface {
	Insert(ctx context.Context, userInfo request.ShopRegister) (*entity.Shop, error)
	GetByUserId(ctx context.Context, userId string) (*entity.Shop, error)
}
