package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shop struct {
	Id        uuid.UUID `json:"id"`
	ShopName  string    `json:"shop_name"`
	Avatar    string    `json:"avatar"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewShop(id uuid.UUID, shopName, avatar, userId string) Shop {
	return Shop{
		Id:        id,
		ShopName:  shopName,
		Avatar:    avatar,
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *Shop) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New()
	return
}
