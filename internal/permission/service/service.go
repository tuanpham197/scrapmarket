package service

import (
	"context"
	"sendo/internal/permission/service/entity"
	"sendo/internal/permission/service/requests"
)

type service struct {
	roleRepo       RoleRepository
	permissionRepo PermissionRepository
}

func NewService(roleRepo RoleRepository, permissionRepo PermissionRepository) service {
	return service{roleRepo: roleRepo, permissionRepo: permissionRepo}
}

func (s service) Create(ctx context.Context, role *requests.RoleRequest) (*entity.Role, error) {
	result, err := s.roleRepo.Create(ctx, role)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s service) GivePermissionTo(ctx context.Context, roleId uint, permissionRequest *requests.AssignPermissionRequest) (bool, error) {

	result, err := s.roleRepo.GivePermissionTo(ctx, roleId, permissionRequest)

	if err != nil {
		return false, err
	}

	return result, nil
}

func (s service) RevokePermissionTo(ctx context.Context, role *entity.Role, permissions entity.Permission) (bool, error) {
	panic("error")
}
