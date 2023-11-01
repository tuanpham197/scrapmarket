package entity

import (
	"time"
)

type Role struct {
	ID          uint          `json:"id" gorm:"primaryKey"`
	Name        string        `json:"name"`
	GuardName   string        `json:"guard_name"`
	Permissions *[]Permission `json:"permissions,omitempty" gorm:"many2many:role_has_permissions;"`
	CreatedAt   *time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

type Permission struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	GuardName string     `json:"guard_name"`
	Role      *[]Role    `json:"roles" gorm:"many2many:role_has_permissions;"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
