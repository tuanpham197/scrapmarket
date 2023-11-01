package mysql

import (
	"context"
	"errors"
	"sendo/internal/shop_category/service"
	"sendo/internal/shop_category/service/entity"
	"sendo/internal/shop_category/service/request"

	productEntity "sendo/internal/product/service/entity"

	"gorm.io/gorm"
)

type shopCategoryRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) service.ShopCategoryRepository {
	return shopCategoryRepo{db: db}
}

func (repo shopCategoryRepo) Insert(ctx context.Context, createShopCategory request.CreateShopCategory) (*entity.ShopCategory, error) {
	tx := repo.db.Begin().Debug()

	defer tx.Commit()

	var products []productEntity.Product
	productListError := tx.Where("id IN (?) AND shop_id = ?", createShopCategory.ProductIds, createShopCategory.ShopId).Find(&products).Error
	if productListError != nil {
		tx.Rollback()
		return nil, nil
	}

	// Check product of shop
	if len(products) != len(createShopCategory.ProductIds) || len(createShopCategory.ProductIds) == 0 {
		var productError = errors.New("product ids not match")
		return nil, productError
	}

	shopCategory := entity.ShopCategory{
		Name:     createShopCategory.Name,
		ShopId:   createShopCategory.ShopId,
		Products: &products,
	}

	result := tx.Omit("Products.*").Create(&shopCategory).Error
	if result != nil {
		tx.Rollback()
		return nil, result
	}

	return &shopCategory, nil
}

func (repo shopCategoryRepo) UpdateShopCategory(ctx context.Context, updateShopCategory request.UpdateShopCategory, shopId string, categoryId string) (*entity.ShopCategory, error) {
	tx := repo.db.Begin().Debug()

	defer tx.Commit()

	// Get ShopCategory by id
	var shopCategory entity.ShopCategory
	err := tx.Preload("Products").Where("id = ?", categoryId).First(&shopCategory).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var products []productEntity.Product
	tx.Where("id IN (?) AND shop_id = ?", updateShopCategory.ProductIds, shopId).Find(&products)

	if len(products) != len(updateShopCategory.ProductIds) || len(updateShopCategory.ProductIds) == 0 {
		var productError = errors.New("product ids not match")
		return nil, productError
	}

	// Update New ShopCategoryProduct
	errAppend := tx.Model(&shopCategory).Omit("Products.*").Association("Products").Replace(&products)
	if errAppend != nil {
		tx.Rollback()
		return nil, errAppend
	}

	return &shopCategory, nil
}

func (repo shopCategoryRepo) DeleteShopCategory(ctx context.Context, shopId string, categoryId string) (*entity.ShopCategory, error) {
	repo.db.Debug()

	// // Get ShopCategory by id
	var shopCategory entity.ShopCategory
	err := repo.db.Where("id = ?", categoryId).First(&shopCategory).Error
	if err != nil {
		return nil, err
	}

	// check ShopId
	if shopCategory.ShopId != shopId {
		var shopIdError = errors.New("can not delete")
		return nil, shopIdError
	}

	// Update New ShopCategoryProduct
	errDelete := repo.db.Delete(&entity.ShopCategory{}, "id = ?", categoryId).Error
	if errDelete != nil {
		return nil, errDelete
	}

	return nil, nil
}

func (repo shopCategoryRepo) GetListShopCategoryAndProduct(ctx context.Context, shopId string) (*[]entity.ShopCategory, error) {
	var shopCategories []entity.ShopCategory
	err := repo.db.Preload("Products").Where("shop_id = ?", shopId).Find(&shopCategories).Error
	if err != nil {
		return nil, err
	}
	return &shopCategories, nil
}
