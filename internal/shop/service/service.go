package service

import (
	"context"
	"sendo/internal/shop/service/entity"
	"sendo/internal/shop/service/request"
)

type shopService struct {
	shopRepository ShopRepository
}

func NewShopService(shopRepository ShopRepository) shopService {
	return shopService{shopRepository: shopRepository}
}

// Register      Register
// @Summary      Register new shop
// @Description  Register new shop
// @Param 		 request body request.ShopRegister true "register param"
// @Tags         shop
// @Produce      json
// @Success		 200	{object} request.ShopRegisterResponse
// @Failure		 400	{object} error
// @Router       /shops/register [post]
func (s shopService) ShopRegister(ctx context.Context, shopRegister request.ShopRegister) (*request.ShopRegisterResponse, error) {

	// check shop exists
	shop, _ := s.shopRepository.GetByUserId(ctx, shopRegister.UserId)

	if shop != nil {
		return nil, request.ShopExistsError{}
	}

	// Insert user
	result, err := s.shopRepository.Insert(ctx, shopRegister)
	if err != nil {
		return nil, err
	}

	return &request.ShopRegisterResponse{
		ShopName: result.ShopName,
		Avatar:   result.Avatar,
		UserId:   result.UserId,
	}, nil
}

func (s shopService) GetShopInfo(ctx context.Context, userId string) (*entity.Shop, error) {
	shop, err := s.shopRepository.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return shop, nil
}
