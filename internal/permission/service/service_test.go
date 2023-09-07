package service

import (
	"context"
	"sendo/internal/permission/service/entity"
	"sendo/internal/permission/service/requests"
	"testing"
)

// Repare data test
var ()

type mockRoleRepository struct{}
type mockPermissionRepository struct{}

func (mockRoleRepository) Create(ctx context.Context, role *requests.RoleRequest) (*entity.Role, error) {
	return nil, nil
}

func (mockRoleRepository) GetOne(ctx context.Context, roleId uint) (*entity.Role, error) {
	return nil, nil
}

func (mockRoleRepository) RevokePermissionTo(ctx context.Context, role *entity.Role, permissions entity.Permission) (bool, error) {
	return false, nil
}

func (mockPermissionRepository) CreatePermission(ctx context.Context, name, guard_name string) (*entity.Permission, error) {
	return nil, nil
}

func TestService_Create(t *testing.T) {

}

func TestService_GivePermissionTo(t *testing.T) {

}

func TestService_RevokePermissionTo(t *testing.T) {

}
