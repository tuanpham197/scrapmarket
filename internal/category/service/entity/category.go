package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type Category struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Thumbnail string     `json:"thumbnail"`
	ParentID  *uuid.UUID `json:"parent_id"`
	Parent    *Category  `gorm:"foreignkey:ParentID" json:"parent,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	c.Id = uuid.New()
	return
}

func (c *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}

func NewCategory(name, thumbnail string, parent_id, shop_id *uuid.UUID) Category {
	return Category{
		Name:      name,
		ParentID:  parent_id,
		Thumbnail: thumbnail,
	}
}
