package service

import (
	"context"
	"sendo/internal/product/service/entity"
	"sendo/internal/product/service/request"
)

type VariantUseCase interface {
	CreateVariant(ctx context.Context, createVariant request.CreateVariant) (*request.CreateVariantResponse, error)
}

type VariantRepository interface {
	InsertVariant(ctx context.Context, createVariant request.CreateVariant) (*entity.Variant, error)
}
