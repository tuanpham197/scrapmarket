package service

import (
	"context"
	"sendo/internal/product/service/entity"
	"sendo/internal/product/service/request"
)

type ConfigUseCase interface {
	CreateConfig(ctx context.Context, createConfig request.CreateConfig) (*request.CreateConfigResponse, error)
}

type ConfigRepository interface {
	InsertConfig(ctx context.Context, createConfig request.CreateConfig) (*entity.Config, error)
	GetOne(ctx context.Context, req request.FormRequest) (*entity.Config, error)
}
