package request

import (
	"sendo/internal/shop/service/entity"

	"github.com/google/uuid"
)

type UserInfo struct {
	Id        uuid.UUID    `json:"id"`
	UserName  string       `json:"username"`
	Email     string       `json:"email"`
	Birthday  string       `json:"birthday"`
	Age       int16        `json:"age"`
	LastName  string       `json:"last_name"`
	FirstName string       `json:"first_name"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	Shop      *entity.Shop `json:"shop"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	UserName  string `json:"username" binding:"required"`
	LastName  string `json:"last_name" binding:"omitempty"`
	FirstName string `json:"first_name" binding:"omitempty"`
	Email     string `json:"email" binding:"email"`
	Salt      string `json:"salt" binding:"omitempty"`
	Password  string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	UserInfo     UserInfo `json:"user_info"`
}

type UserExistsError struct{}

func (e UserExistsError) Error() string {
	return "User already exists"
}

type ResponseMessageError struct {
	Message string `json:"message"`
}

func (e ResponseMessageError) Error() string {
	return e.Message
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
