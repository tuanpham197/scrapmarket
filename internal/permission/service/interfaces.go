package service

import (
	"context"
	"sendo/internal/permission/service/entity"
	"sendo/internal/permission/service/requests"
)

type RoleUseCase interface {
	Create(ctx context.Context, role *requests.RoleRequest) (*entity.Role, error)
	GivePermissionTo(ctx context.Context, roleId uint, permissionRequest *requests.AssignPermissionRequest) (bool, error)
	RevokePermissionTo(ctx context.Context, role *entity.Role, permissions entity.Permission) (bool, error)
}

type RoleRepository interface {
	Create(ctx context.Context, role *requests.RoleRequest) (*entity.Role, error)
	GivePermissionTo(ctx context.Context, roleId uint, permissionRequest *requests.AssignPermissionRequest) (bool, error)
	GetOne(ctx context.Context, roleId uint) (*entity.Role, error)
}

type Permission interface {
	CreatePermission(ctx context.Context, name, guard_name string) (*entity.Permission, error)
}

type PermissionRepository interface {
	CreatePermission(ctx context.Context, name, guard_name string) (*entity.Permission, error)
}
