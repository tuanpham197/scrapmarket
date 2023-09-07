package service

import (
	"context"
	"sendo/internal/product/service/request"
)

type variantService struct {
	variantRepository VariantRepository
}

func NewVariantService(variantRepository VariantRepository) variantService {
	return variantService{variantRepository: variantRepository}
}

// Create      Create
// @Summary      Create variant
// @Description  Create variant
// @Param 		 request body request.CreateVariant true "create param"
// @Tags         product
// @Produce      json
// @Success		 200	{object} request.CreateVariantResponse
// @Failure		 400	{object} error
// @Router       /variants/create [post]
func (v variantService) CreateVariant(ctx context.Context, createVariant request.CreateVariant) (*request.CreateVariantResponse, error) {

	// Insert variant
	result, err := v.variantRepository.InsertVariant(ctx, createVariant)
	if err != nil {
		return nil, err
	}

	return &request.CreateVariantResponse{
		ProductId:    result.ProductId,
		ConfigValue1: *result.ConfigValue1,
		ConfigValue2: *result.ConfigValue2,
		Price:        result.Price,
		SalePrice:    result.SalePrice,
		Quantity:     result.Quantity,
	}, nil
}
