package mysql

import (
	"context"
	"sendo/internal/product/service"
	"sendo/internal/product/service/entity"
	"sendo/internal/product/service/request"

	"gorm.io/gorm"
)

type configValueRepo struct {
	db *gorm.DB
}

func NewConfigValueRepo(db *gorm.DB) service.ConfigValueRepository {
	return configValueRepo{db: db}
}

func (repo configValueRepo) InsertConfigValue(ctx context.Context, createConfigValue request.CreateConfigValue) (*entity.ConfigValue, error) {
	configValue := entity.ConfigValue{
		Name:     createConfigValue.Name,
		ConfigId: createConfigValue.ConfigId,
	}

	result := repo.db.Create(&configValue)
	if result.Error != nil {
		return nil, result.Error
	}
	return &configValue, nil
}
