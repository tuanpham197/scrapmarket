package mysql

import (
	"context"
	"sendo/internal/product/service"
	"sendo/internal/product/service/entity"
	"sendo/internal/product/service/request"

	"gorm.io/gorm"
)

type configRepo struct {
	db *gorm.DB
}

func (repo configRepo) GetOne(ctx context.Context, req request.FormRequest) (*entity.Config, error) {
	//TODO implement me
	panic("implement me")
}

func NewConfigRepo(db *gorm.DB) service.ConfigRepository {
	return configRepo{db: db}
}

func (repo configRepo) InsertConfig(ctx context.Context, createConfig request.CreateConfig) (*entity.Config, error) {

	config := entity.Config{
		Name:   createConfig.Name,
		ShopId: createConfig.ShopId,
	}

	result := repo.db.Create(&config)
	if result.Error != nil {
		return nil, result.Error
	}
	return &config, nil
}
