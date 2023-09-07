package request

type CreateProduct struct {
	Name        string  `json:"name" binding:"required"`
	ShopId      string  `json:"shop_id" binding:"required"`
	ThumbNail   string  `json:"thumbnail" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	SalePrice   float32 `json:"sale_price"`
	Description string  `json:"description"`
}

type CreateProductResponse struct {
	Name        string  `json:"name"`
	ShopId      string  `json:"shop_id"`
	ThumbNail   string  `json:"thumbnail"`
	Price       float32 `json:"price"`
	SalePrice   float32 `json:"sale_price"`
	Description string  `json:"description"`
}

type CreateConfigValueRaw struct {
	Name string `json:"name"`
}

type CreateConfigRaw struct {
	Name         string                 `json:"name"`
	ConfigValues []CreateConfigValueRaw `json:"config_values"`
}

type CreateVariantRaw struct {
	ConfigValue1 string  `json:"config_value_1"`
	ConfigValue2 string  `json:"config_value_2"`
	Price        float32 `json:"price"`
	Quantity     int     `json:"quantity"`
}

type CreateProductRaw struct {
	Name        string             `json:"name" binding:"required"`
	Thumbnail   string             `json:"thumbnail" binding:"required"`
	CategoryId  string             `json:"category_id" binding:"required"`
	Description string             `json:"description"`
	Price       float32            `json:"price"`
	SalePrice   float32            `json:"sale_price"`
	Quantity    int                `json:"quantity"`
	Configs     []CreateConfigRaw  `json:"configs"`
	Variants    []CreateVariantRaw `json:"variants"`
}

type FilterRequest struct {
	// ID         string  `uri:"id" binding:"required,uuid"`
	CategoryId string  `form:"category_id"`
	ShopId     string  `form:"shop_id"`
	Name       string  `form:"name"`
	PriceFrom  float32 `form:"price_from"`
	PriceTo    float32 `form:"price_to"`
	Page       int     `form:"page"`
	PerPage    int     `form:"per_page" default:"10"`
}

type IDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}
