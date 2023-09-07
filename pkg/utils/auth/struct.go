package auth_util

import (
	"sendo/internal/permission/service/entity"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	UserID uuid.UUID      `json:"user_id"`
	ShopID *uuid.UUID     `json:"shop_id"`
	Roles  *[]entity.Role `json:"roles"`
}

type CustomClaims struct {
	UserID uuid.UUID      `json:"userId"`
	ShopID *uuid.UUID     `json:"shopId"`
	Roles  *[]entity.Role `json:"roles"`
	jwt.RegisteredClaims
}
