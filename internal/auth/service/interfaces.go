package service

import (
	"context"
	"sendo/internal/auth/service/entity"
	"sendo/internal/auth/service/request"
)

type UserUseCase interface {
	// FetchUser(ctx context.Context) ([]entity.User, error)
	Login(ctx context.Context, userLogin request.UserLogin) (*request.UserLoginResponse, error)
	Register(ctx context.Context, userRegister request.UserRegister) (*request.UserLoginResponse, error)
	RefreshToken(ctx context.Context, token request.RefreshTokenRequest) (*request.UserLoginResponse, error)
	GetInfoUser(ctx context.Context, id string) (*entity.User, error)
	AssignRoleUser(ctx context.Context, roleId request.AssignRoleUser, userId string) (bool, error)
	LoginByGoogle(ctx context.Context, email string) (*entity.User, error)
}

type UserRepository interface {
	Insert(ctx context.Context, userInfo request.UserRegister) (*entity.User, error)
	GetOne(ctx context.Context, id string) (*entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Delete(ctx context.Context, id string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	AssignRoleUser(ctx context.Context, roleId request.AssignRoleUser, userId string) (bool, error)
}
