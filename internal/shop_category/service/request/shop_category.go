package request

import "github.com/google/uuid"

type CreateShopCategory struct {
	Name       string      `json:"name" binding:"required"`
	ShopId     string      `json:"shop_id"`
	ProductIds []uuid.UUID `json:"product_ids"`
}

type CreateShopCategoryResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	ShopId string `json:"shop_id"`
}

type UpdateShopCategory struct {
	ProductIds []uuid.UUID `json:"product_ids"`
}
