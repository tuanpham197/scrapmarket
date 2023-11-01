package request

type CreateConfig struct {
	Name   string `json:"name" binding:"required"`
	ShopId string `json:"shop_id" binding:"required"`
}

type CreateConfigResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShopId    string `json:"shop_id"`
	ProductId string `json:"product_id"`
}

type FormRequest struct {
	ProductId string `form:"product_id"`
}
