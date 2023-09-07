package entity

import (
	"errors"
	"time"

	"sendo/internal/product/service/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type ShopCategory struct {
	Id        uuid.UUID         `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	Name      string            `json:"name"`
	ShopId    string            `json:"shop_id"`
	Products  *[]entity.Product `json:"products" gorm:"many2many:shop_category_products;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func NewShopCategory(id uuid.UUID, name, shop_id string) ShopCategory {
	return ShopCategory{
		Id:        id,
		Name:      name,
		ShopId:    shop_id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (ShopCategory) TableName() string {
	return "shop_categories"
}

func (s *ShopCategory) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New()
	return
}
