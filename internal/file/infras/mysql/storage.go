package mysql

import (
	"sendo/internal/file/service"

	"gorm.io/gorm"
)

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) service.FileRepository {
	return mysqlRepo{db: db}
}
