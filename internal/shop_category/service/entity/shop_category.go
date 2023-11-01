package entity

import (
	"time"

	"sendo/internal/product/service/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShopCategory struct {
	Id        uuid.UUID         `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	Name      string            `json:"name"`
	ShopId    string            `json:"shop_id"`
	Products  *[]entity.Product `json:"products" gorm:"many2many:shop_category_products;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func NewShopCategory(id uuid.UUID, name, shopId string) ShopCategory {
	return ShopCategory{
		Id:        id,
		Name:      name,
		ShopId:    shopId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *ShopCategory) TableName() string {
	return "shop_categories"
}

func (s *ShopCategory) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New()
	return
}
