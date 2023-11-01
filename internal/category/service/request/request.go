package request

type CategoryCreateRequest struct {
	Name      string `json:"name" binding:"required"`
	Thumbnail string `json:"thumbnail" binding:"omitempty"`
	ParentId  string `json:"parent_id" binding:"omitempty"`
	Type      int    `json:"type" binding:"omitempty"`
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
