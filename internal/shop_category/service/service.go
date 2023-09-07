package service

import (
	"context"
	"sendo/internal/shop_category/service/entity"
	"sendo/internal/shop_category/service/request"
)

type shopCategoryService struct {
	shopCategoryRepository ShopCategoryRepository
}

func NewShopService(shopCategoryRepository ShopCategoryRepository) shopCategoryService {
	return shopCategoryService{shopCategoryRepository: shopCategoryRepository}
}

// CreateShopCategory      CreateShopCategory
// @Summary      CreateShopCategory
// @Description  CreateShopCategory
// @Param 		 request body request.CreateShopCategory true "register param"
// @Tags         Shop category
// @Produce      json
// @Success		 200	{object} request.CreateShopCategoryResponse
// @Failure		 400	{object} error
// @Router       /shop-categories [post]
func (s shopCategoryService) CreateShopCategory(ctx context.Context, createShopCategory request.CreateShopCategory) (*request.CreateShopCategoryResponse, error) {

	result, err := s.shopCategoryRepository.Insert(ctx, createShopCategory)
	if err != nil {
		return nil, err
	}

	return &request.CreateShopCategoryResponse{
		Id:     result.Id.String(),
		Name:   result.Name,
		ShopId: result.ShopId,
	}, nil
}

// UpdateShopCategory      UpdateShopCategory
// @Summary      UpdateShopCategory
// @Description  UpdateShopCategory
// @Param 		 id path string  true  "ID"
// @Param 		 request body request.UpdateShopCategory true "register param"
// @Tags         Shop category
// @Produce      json
// @Success		 200	{object} entity.ShopCategory
// @Failure		 400	{object} error
// @Router       /shop-categories/{id} [patch]
func (s shopCategoryService) UpdateShopCategory(ctx context.Context, updateShopCategory request.UpdateShopCategory, shopId string, categoryId string) (*entity.ShopCategory, error) {

	result, err := s.shopCategoryRepository.UpdateShopCategory(ctx, updateShopCategory, shopId, categoryId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete        Delete
// @Summary      Delete shop category
// @Description  Delete shop category
// @Param 		 id path string true "Shop category id"
// @Tags         Shop category
// @Produce      json
// @Success		 200	{object} entity.ShopCategory
// @Failure		 400	{object} error
// @Router       /shop-categories/:id [delete]
func (s shopCategoryService) DeleteShopCategory(ctx context.Context, shopId string, categoryId string) (*entity.ShopCategory, error) {

	result, err := s.shopCategoryRepository.DeleteShopCategory(ctx, shopId, categoryId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get list      List
// @Summary      Get list shop category and product
// @Description  Get list shop category and product
// @Param 		 id path string true "Shop category id"
// @Tags         Shop category
// @Produce      json
// @Success		 200	{object} []entity.ShopCategory
// @Failure		 400	{object} error
// @Router       /shop-categories/:id [get]
func (s shopCategoryService) GetListShopCategoryAndProduct(ctx context.Context, shopId string) (*[]entity.ShopCategory, error) {

	result, err := s.shopCategoryRepository.GetListShopCategoryAndProduct(ctx, shopId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
