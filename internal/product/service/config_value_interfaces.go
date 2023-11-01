package service

import (
	"context"
	"sendo/internal/product/service/entity"
	"sendo/internal/product/service/request"
)

type ConfigValueUseCase interface {
	CreateConfigValue(ctx context.Context, createConfigValue request.CreateConfigValue) (*request.CreateConfigValueResponse, error)
}

type ConfigValueRepository interface {
	InsertConfigValue(ctx context.Context, createConfigValue request.CreateConfigValue) (*entity.ConfigValue, error)
}
