package service

import (
	"context"
	"sendo/internal/product/service/entity"
	"sendo/internal/product/service/request"
	"sendo/pkg/utils/paginations"
)

type productService struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) productService {
	return productService{productRepository: productRepository}
}

// Create      Create
// @Summary      Create product
// @Description  Create product
// @Param 		 request body request.CreateProduct true "create param"
// @Tags         product
// @Produce      json
// @Success		 200	{object} request.CreateProductResponse
// @Failure		 400	{object} error
// @Router       /products/create [post]
func (p productService) CreateProduct(ctx context.Context, createProduct request.CreateProductRaw, shopId string) (*request.CreateProductResponse, error) {

	// Insert product
	result, err := p.productRepository.InsertProduct(ctx, createProduct, shopId)
	if err != nil {
		return nil, err
	}

	return &request.CreateProductResponse{
		Name:        result.Name,
		ShopId:      result.ShopId,
		ThumbNail:   result.ThumbNail,
		Price:       result.Price,
		SalePrice:   result.SalePrice,
		Description: result.Description,
	}, nil
}

// Get list by catetory      Get list by catetory
// @Summary      Get product
// @Description  Get product
// @Param 		 request body string true "get param"
// @Tags         product
// @Produce      json
// @Success		 200	{object} paginations.Pagination
// @Failure		 400	{object} error
// @Router       /products/categories/:id [get]
func (p productService) GetList(ctx context.Context, filter *request.FilterRequest) (*paginations.Pagination, error) {
	result, err := p.productRepository.GetList(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p productService) GetDetail(ctx context.Context, id string) (*entity.Product, error) {
	product, err := p.productRepository.GetDetail(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p productService) GetConfigProduct(ctx context.Context, productId string) (*[]entity.Config, error) {
	product, err := p.productRepository.GetConfigProduct(ctx, productId)
	if err != nil {
		return nil, err
	}

	return product, nil
}
