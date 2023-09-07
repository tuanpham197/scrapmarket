package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type Shop struct {
	Id        uuid.UUID `json:"id"`
	ShopName  string    `json:"shop_name"`
	Avatar    string    `json:"avatar"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewShop(id uuid.UUID, shop_name, avatar, user_id string) Shop {
	return Shop{
		Id:        id,
		ShopName:  shop_name,
		Avatar:    avatar,
		UserId:    user_id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *Shop) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New()
	return
}
