package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConfigValue struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ConfigId  string    `json:"config_id"`
	Config    Config    `json:"config"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewConfigValue(id uuid.UUID, name, configId string) ConfigValue {
	return ConfigValue{
		Id:        id,
		Name:      name,
		ConfigId:  configId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *ConfigValue) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New()
	return
}
