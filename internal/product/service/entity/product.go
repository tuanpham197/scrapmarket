package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          uuid.UUID `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	Name        string    `json:"name"`
	ShopId      string    `json:"shop_id"`
	CategoryId  string    `json:"category_id"`
	ThumbNail   string    `json:"thumbnail" gorm:"column:thumbnail"`
	Price       float32   `json:"price"`
	SalePrice   float32   `json:"sale_price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewProduct(id uuid.UUID, name, shopId, thumbnail, description string, price, salePrice float32) Product {
	return Product{
		Id:          id,
		Name:        name,
		ShopId:      shopId,
		ThumbNail:   thumbnail,
		Price:       price,
		SalePrice:   salePrice,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (Product) TableName() string {
	return "products"
}

func (s *Product) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New()
	return
}
