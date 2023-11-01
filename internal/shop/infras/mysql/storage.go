package mysql

import (
	"context"
	"sendo/internal/shop/service"
	"sendo/internal/shop/service/entity"
	"sendo/internal/shop/service/request"

	"gorm.io/gorm"
)

type shopRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) service.ShopRepository {
	return shopRepo{db: db}
}

func (repo shopRepo) Insert(ctx context.Context, shopRegister request.ShopRegister) (*entity.Shop, error) {

	shop := entity.Shop{
		ShopName: shopRegister.ShopName,
		Avatar:   shopRegister.Avatar,
		UserId:   shopRegister.UserId,
	}

	result := repo.db.Create(&shop)
	if result.Error != nil {
		return nil, result.Error
	}
	return &shop, nil
}

func (repo shopRepo) GetByUserId(ctx context.Context, userId string) (*entity.Shop, error) {
	var shop entity.Shop
	err := repo.db.First(&shop, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	return &shop, nil
}
