package requests

type RoleRequest struct {
	Name      string `json:"name" validate:"required"`
	GuardName string `json:"guard_name" validate:"required"`
}

type PermissionRequest struct {
	Name      string `json:"name" validate:"required"`
	GuardName string `json:"guard_name" validate:"required"`
}

type AssignPermissionRequest struct {
	Permission    string `json:"permission" validate:"required_without=PermissionIds"`
	PermissionIds []int  `json:"permission_ids" validate:"required_without=Permission"`
}
