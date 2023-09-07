package mysql

import (
	"context"
	"sendo/internal/product/service"
	"sendo/internal/product/service/entity"
	"sendo/internal/product/service/request"
	"sendo/pkg/utils/paginations"

	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) service.ProductRepository {
	return productRepo{db: db}
}

func (repo productRepo) InsertProduct(ctx context.Context, createProductRaw request.CreateProductRaw, shopId string) (*entity.Product, error) {
	tx := repo.db.Begin()

	product := entity.Product{
		Name:        createProductRaw.Name,
		ThumbNail:   createProductRaw.Thumbnail,
		CategoryId:  createProductRaw.CategoryId,
		ShopId:      shopId,
		Price:       createProductRaw.Price,
		SalePrice:   createProductRaw.SalePrice,
		Description: createProductRaw.Description,
	}

	result := tx.Create(&product).Error

	if result != nil {
		tx.Rollback()
		return nil, result
	}

	// If product doesn't have configs and variants
	if len(createProductRaw.Configs) <= 0 {
		variantObject := entity.Variant{
			ProductId: product.Id.String(),
			Price:     createProductRaw.Price,
			SalePrice: createProductRaw.SalePrice,
			Quantity:  createProductRaw.Quantity,
		}
		variantResult := tx.Create(&variantObject).Error
		if variantResult != nil {
			tx.Rollback()
			return nil, result
		}

		tx.Commit()
		return &product, nil
	}

	mappingConfigValue := make(map[string]string)

	for _, config := range createProductRaw.Configs {

		var configValueArr []entity.ConfigValue

		for _, configValueEntity := range config.ConfigValues {
			configValueArr = append(configValueArr, entity.ConfigValue{Name: configValueEntity.Name})
		}

		configObject := entity.Config{
			Name:         config.Name,
			ShopId:       shopId,
			ConfigValues: configValueArr,
		}

		// Create Config and Config Value
		configResult := tx.Create(&configObject).Error

		if configResult != nil {
			tx.Rollback()
			return nil, result
		}

		for _, conValue := range configValueArr {
			mappingConfigValue[config.Name+"-"+conValue.Name] = conValue.Id.String()
		}

	}

	// create variants
	for _, variant := range createProductRaw.Variants {
		value1 := mappingConfigValue[variant.ConfigValue1]
		value2 := mappingConfigValue[variant.ConfigValue2]
		variantObject := entity.Variant{
			ProductId:    product.Id.String(),
			ConfigValue1: &value1,
			ConfigValue2: &value2,
			Price:        variant.Price,
			Quantity:     variant.Quantity,
		}

		variantResult := tx.Create(&variantObject).Error

		if variantResult != nil {
			tx.Rollback()
			return nil, result
		}
	}

	//TODO: update product price
	//TODO: update product images

	tx.Commit()
	return &product, nil
}

func (repo productRepo) GetList(ctx context.Context, filter *request.FilterRequest) (*paginations.Pagination, error) {

	query := repo.db.Model(entity.Product{})
	var products []entity.Product

	if filter.ShopId != "" {
		query.Where("shop_id = ?", filter.ShopId)
	}

	if filter.CategoryId != "" {
		query.Where("category_id = ?", filter.CategoryId)
	}
	if len(filter.Name) > 0 {
		query.Where("name like ?", "%"+filter.Name+"%")
	}

	if filter.PriceFrom != 0 {
		query.Where("price >= ?", filter.PriceFrom)
	}

	if filter.PriceTo != 0 {
		query.Where("price <= ?", filter.PriceTo)
	}

	p := &paginations.Param{
		DB:      repo.db,
		Query:   query,
		Page:    filter.Page,
		Limit:   filter.PerPage,
		OrderBy: []string{},
		ShowSQL: false,
	}
	result, err := paginations.Pagging(p, &products)
	if err != nil {
		return nil, err
	}
	return result, nil

}
