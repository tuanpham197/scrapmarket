package mysql

import (
	"context"
	"sendo/internal/product/service"
	"sendo/internal/product/service/entity"
	"sendo/internal/product/service/request"

	"gorm.io/gorm"
)

type variantRepo struct {
	db *gorm.DB
}

func NewVariantRepo(db *gorm.DB) service.VariantRepository {
	return variantRepo{db: db}
}

func (repo variantRepo) InsertVariant(ctx context.Context, createVariant request.CreateVariant) (*entity.Variant, error) {

	variant := entity.Variant{
		ProductId:    createVariant.ProductId,
		ConfigValue1: &createVariant.ConfigValue1,
		ConfigValue2: &createVariant.ConfigValue2,
		Price:        createVariant.Price,
		SalePrice:    createVariant.SalePrice,
		Quantity:     createVariant.Quantity,
	}

	result := repo.db.Create(&variant)
	if result.Error != nil {
		return nil, result.Error
	}
	return &variant, nil
}
