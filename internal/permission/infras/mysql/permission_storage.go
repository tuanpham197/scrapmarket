package mysql

import (
	"context"
	"sendo/internal/permission/service"
	"sendo/internal/permission/service/entity"

	"gorm.io/gorm"
)

func NewMySQLRepoPermission(db *gorm.DB) service.PermissionRepository {
	return mysqlRepo{db: db}
}

func (repo mysqlRepo) CreatePermission(ctx context.Context, name, guard_name string) (*entity.Permission, error) {
	permission := entity.Permission{Name: name, GuardName: guard_name}
	result := repo.db.Create(&permission)
	if result.Error != nil {
		return nil, result.Error
	}

	return &permission, nil
}
