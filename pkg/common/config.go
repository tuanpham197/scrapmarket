package common

import "gorm.io/gorm"

type GormComponent interface {
	GetDB() *gorm.DB
}

type Config interface {
	GetGRPCPort() int
	GetGRPCServerAddress() string
}
