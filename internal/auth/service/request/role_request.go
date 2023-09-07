package request

type AssignRoleUser struct {
	Roles []uint `json:"roles" binding:"required"`
}

type UserRoleRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}
