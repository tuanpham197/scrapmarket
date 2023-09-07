package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type Config struct {
	Id           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	ShopId       string        `json:"shop_id"`
	ConfigValues []ConfigValue `json:"config_values"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

func NewConfig(id uuid.UUID, name, shopId string) Config {
	return Config{
		Id:        id,
		Name:      name,
		ShopId:    shopId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *Config) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New()
	return
}
