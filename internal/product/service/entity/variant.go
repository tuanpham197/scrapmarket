package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	Id           uuid.UUID  `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	ProductId    string     `json:"product_id" gorm:"column:product_id"`
	ConfigValue1 *string    `json:"config_value_1,omitempty" gorm:"column:config_value_1"`
	ConfigValue2 *string    `json:"config_value_2,omitempty" gorm:"column:config_value_2"`
	Price        float32    `json:"price" gorm:"column:price"`
	SalePrice    float32    `json:"sale_price" gorm:"column:sale_price"`
	Quantity     int        `json:"quantity" gorm:"column:quantity"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func NewVariant(id uuid.UUID, name, productId, configValue1, configValue2 string, price, salePrice float32, quantity int) Variant {
	timeNow := time.Now()
	return Variant{
		Id:           id,
		ProductId:    productId,
		ConfigValue1: &configValue1,
		ConfigValue2: &configValue2,
		Price:        price,
		SalePrice:    salePrice,
		Quantity:     quantity,
		CreatedAt:    &timeNow,
		UpdatedAt:    &timeNow,
	}
}

func (s *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New()
	return
}
