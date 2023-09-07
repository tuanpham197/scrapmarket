package mysql

import (
	"context"
	"sendo/internal/permission/service"
	"sendo/internal/permission/service/entity"
	"sendo/internal/permission/service/requests"

	"gorm.io/gorm"
)

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) service.RoleRepository {
	return mysqlRepo{db: db}
}

func (repo mysqlRepo) Create(ctx context.Context, role *requests.RoleRequest) (*entity.Role, error) {
	roleModel := entity.Role{Name: role.Name, GuardName: role.GuardName}
	result := repo.db.Create(&roleModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &roleModel, nil
}

func (repo mysqlRepo) GetOne(ctx context.Context, roleId uint) (*entity.Role, error) {
	var role entity.Role

	err := repo.db.First(&role).Where("id = ?", roleId).Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (repo mysqlRepo) GivePermissionTo(ctx context.Context, roleId uint, permissionRequest *requests.AssignPermissionRequest) (bool, error) {
	tx := repo.db

	// get role by id
	var role entity.Role
	err := tx.First(&role).Where("id = ?", roleId).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}

	// Initialize it with an empty slice
	role.Permissions = &[]entity.Permission{}

	// check create new permission
	permissionName := permissionRequest.Permission
	if permissionName != "" {
		permission := entity.Permission{
			Name:      permissionName,
			GuardName: "api",
		}

		errCreatePermission := tx.Create(&permission).Error
		if errCreatePermission != nil {
			tx.Rollback()
			return false, errCreatePermission
		}

		// assing permission for role
		errAppend := tx.Model(&role).Omit("Permissions.*").Association("Permissions").Replace(&permission)
		if errAppend != nil {
			tx.Rollback()
			return false, errAppend
		}

		tx.Commit()
		return true, nil
	}

	// passed list permission id
	// get list permission by ids
	permissionIds := permissionRequest.PermissionIds
	var permissions []entity.Permission
	errListPermission := tx.Find(&permissions, permissionIds).Error
	if errListPermission != nil {
		tx.Rollback()
		return false, errListPermission
	}

	// assing permission for role
	errAppend := tx.Model(&role).Omit("Permissions.*").Association("Permissions").Replace(&permissions)
	if errAppend != nil {
		tx.Rollback()
		return false, errAppend
	}

	tx.Commit()
	return true, nil
}
