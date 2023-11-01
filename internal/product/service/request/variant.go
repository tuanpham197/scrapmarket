package request

type CreateVariant struct {
	ProductId    string  `json:"product_id" binding:"required"`
	ConfigValue1 string  `json:"config_value_1"`
	ConfigValue2 string  `json:"config_value_2"`
	Price        float32 `json:"price" binding:"required"`
	SalePrice    float32 `json:"sale_price"`
	Quantity     int     `json:"quantity" binding:"required"`
}

type CreateVariantResponse struct {
	ProductId    string  `json:"product_id"`
	ConfigValue1 string  `json:"config_value_1"`
	ConfigValue2 string  `json:"config_value_2"`
	Price        float32 `json:"price"`
	SalePrice    float32 `json:"sale_price"`
	Quantity     int     `json:"quantity"`
}
