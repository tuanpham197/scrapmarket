package auth_util

import (
	"os"
	"sendo/internal/auth/service/request"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenExpireDuration  = time.Minute * 60 // Access token expiration duration
	refreshTokenExpireDuration = time.Hour * 24   // Refresh token expiration duration
)

var (
	SECRET_KEY = os.Getenv("SECRET_KEY")
)

// Generate an access token
func GenerateAccessToken(payload *Payload) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(accessTokenExpireDuration))

	claims := CustomClaims{
		UserID: payload.UserID,
		ShopID: payload.ShopID,
		Roles:  payload.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_KEY))
}

// Generate a refresh token
func GenerateRefreshToken(payload *Payload) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(refreshTokenExpireDuration))

	claims := CustomClaims{
		UserID: payload.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyRefreshToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, request.ResponseMessageError{
				Message: "Verify refresh token fail",
			}
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, request.ResponseMessageError{
		Message: "Verify refresh token fail",
	}
}
