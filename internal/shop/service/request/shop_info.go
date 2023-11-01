package request

type ShopRegister struct {
	ShopName string `json:"shop_name" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	UserId   string `json:"user_id"`
}

type ShopRegisterResponse struct {
	ShopName string `json:"shop_name"`
	Avatar   string `json:"avatar"`
	UserId   string `json:"user_id"`
}

type ShopExistsError struct{}

func (e ShopExistsError) Error() string {
	return "Shop already exists"
}
