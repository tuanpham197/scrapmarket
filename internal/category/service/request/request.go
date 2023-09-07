package request

type CategoryCreateRequest struct {
	Name      string `form:"name" binding:"required"`
	Thumbnail string `form:"thumbnail" binding:"omitempty"`
	ParentId  string `form:"parent_id" binding:"omitempty"`
	ShopId    string `form:"shop_id" binding:"omitempty"`
	Type      int    `form:"type" binding:"omitempty"`
}

type QueryParam struct {
	Name     string `form:"name"`
	ParentId string `form:"parent_id"`
	Page     int    `form:"page"`
	PerPage  int    `form:"per_page" default:"10"`
}

type RequestURI struct {
	ID string `uri:"id" binding:"required,uuid"`
}
