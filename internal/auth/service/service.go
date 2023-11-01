package service

import (
	"context"
	"errors"
	"sendo/internal/auth/service/entity"
	"sendo/internal/auth/service/request"
	"sendo/pkg/constants"
	authUtil "sendo/pkg/utils/auth"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Hasher interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashedPassword, salt, password string) bool
}

type Service struct {
	repository UserRepository
	hasher     Hasher
	log        *zap.SugaredLogger
}

func NewService(repository UserRepository, hasher Hasher, logger *zap.SugaredLogger) Service {
	//errEnv := godotenv.Load("/home/nlcpu0203/source/me/sendo/scrapmarketbe/.env")
	//if errEnv != nil {
	//	log.Fatal("Error loading .env file 3")
	//}
	return Service{
		repository,
		hasher,
		logger,
	}
}

// Login
// @Summary      Login user
// @Description  Login user
// @Param 		 request body request.UserLogin true "login param"
// @Tags         Auth
// @Produce      json
// @Success		 200	{object} request.UserLoginResponse
// @Failure		 400	{object} error
// @Router       /auth/login [post]
func (s Service) Login(ctx context.Context, userLogin request.UserLogin) (*request.UserLoginResponse, error) {

	user, err := s.repository.GetByEmail(ctx, userLogin.Email)

	if err != nil {
		return nil, constants.ErrUserNotExist
	}

	// TODO: Slow password compare by bcrypt
	resultHash := s.hasher.CompareHashPassword(user.Password, user.Salt, userLogin.Password)
	if !resultHash {
		return nil, constants.ErrWrongPass
	}

	var shopId *uuid.UUID
	if user.Shop != nil {
		shopId = &user.Shop.Id
	}

	// generate access token
	// Assuming the login is successful, generate the tokens
	payload := authUtil.Payload{
		UserID: user.Id,
		ShopID: shopId,
		Roles:  user.Roles,
	}
	accessToken, err := authUtil.GenerateAccessToken(&payload)
	if err != nil {
		return nil, nil
	}
	refreshToken, err := authUtil.GenerateRefreshToken(&payload)
	if err != nil {
		return nil, nil
	}

	// Return the tokens in the response
	return &request.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserInfo: request.UserInfo{
			Id:        user.Id,
			UserName:  user.UserName,
			LastName:  user.LastName,
			Email:     user.Email,
			FirstName: user.FirstName,
			Shop:      user.Shop,
		},
	}, nil

}

// Register      Register
// @Summary      Register new user
// @Description  Register new user
// @Param 		 request body request.UserRegister true "register param"
// @Tags         Auth
// @Produce      json
// @Success		 200	{object} request.UserLoginResponse
// @Failure		 400	{object} error
// @Router       /register [post]
func (s Service) Register(ctx context.Context, userRegister request.UserRegister) (*request.UserLoginResponse, error) {

	// check email exists
	user, _ := s.repository.GetByEmail(ctx, userRegister.Email)
	if user != nil {
		return nil, request.UserExistsError{}
	}

	//generate salt
	salt, errSalt := s.hasher.RandomStr(5)
	if errSalt != nil {
		return nil, errors.New("Generate salt fail")
	}

	// hash pass after call repo
	//
	hashPass, errHash := s.hasher.HashPassword(salt, userRegister.Password)
	if errHash != nil {
		return nil, errors.New("Hash password fail")
	}

	// inser to db
	userRegister.Password = hashPass
	userRegister.Salt = salt
	result, err := s.repository.Insert(ctx, userRegister)
	if err != nil {
		return nil, err
	}

	// Assuming the login is successful, generate the tokens
	payload := authUtil.Payload{
		UserID: result.Id,
	}
	accessToken, err := authUtil.GenerateAccessToken(&payload)
	if err != nil {
		return nil, nil
	}

	refreshToken, err := authUtil.GenerateRefreshToken(&payload)
	if err != nil {
		return nil, nil
	}

	return &request.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserInfo: request.UserInfo{
			Id:        result.Id,
			UserName:  result.UserName,
			LastName:  result.LastName,
			Email:     result.Email,
			FirstName: result.FirstName,
		},
	}, nil
}

// RefreshToken
// @Summary      Refresh token
// @Description  Refresh token
// @Param 		 request body request.RefreshTokenRequest true "token param"
// @Tags         Auth
// @Produce      json
// @Success		 200	{object} request.UserLoginResponse
// @Failure		 400	{object} error
// @Router       /refresh-token [post]
func (s Service) RefreshToken(ctx context.Context, token request.RefreshTokenRequest) (*request.UserLoginResponse, error) {
	// Verify the refresh token and extract the user ID
	claims, err := authUtil.VerifyRefreshToken(token.RefreshToken)
	if err != nil {
		return nil, request.ResponseMessageError{
			Message: "Failed to verify token",
		}
	}

	// Generate a new access token
	payload := authUtil.Payload{
		UserID: claims.UserID,
		ShopID: claims.ShopID,
	}
	accessToken, err := authUtil.GenerateAccessToken(&payload)
	if err != nil {
		return nil, request.ResponseMessageError{
			Message: "Failed to generate access token",
		}
	}
	// Generate new refresh token
	refreshToken, err := authUtil.GenerateRefreshToken(&payload)
	if err != nil {
		return nil, request.ResponseMessageError{
			Message: "Failed to generate refresh token",
		}
	}

	userInfo, err := s.repository.GetOne(ctx, claims.UserID.String())
	if err != nil {
		return nil, constants.ErrUserNotExist
	}

	//TODO: handle revoke old acccess token

	// Return the new access token in the response
	return &request.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserInfo: request.UserInfo{
			Id:        userInfo.Id,
			UserName:  userInfo.UserName,
			LastName:  userInfo.LastName,
			Email:     userInfo.Email,
			FirstName: userInfo.FirstName,
		},
	}, nil
}

// GetInfoUser
// @Summary      Get info
// @Description  Get info
// @Tags         User
// @Produce      json
// @Success		 200	{object} entity.User
// @Failure		 400	{object} error
// @Router       /users/info [post]
func (s Service) GetInfoUser(ctx context.Context, id string) (*entity.User, error) {
	user, err := s.repository.GetOne(ctx, id)
	s.log.Info("Logger", zap.Any("user", user))
	if err != nil {
		return nil, err
	}

	return user, nil
}

// AssignRoleUser
// @Summary      Assign role for user
// @Description  Assign role for user
// @Param 		 request body request.AssignRoleUser true "login param"
// @Param 		 id path string  true  "User ID"
// @Tags         User
// @Produce      json
// @Success		 200	{object} bool
// @Failure		 400	{object} error
// @Router       /:id/assign-role [post]
func (s Service) AssignRoleUser(ctx context.Context, roles request.AssignRoleUser, userId string) (bool, error) {
	result, err := s.repository.AssignRoleUser(ctx, roles, userId)

	if err != nil {
		return false, err
	}

	return result, nil
}

// LoginByGoogle
// @Summary      Login with google
// @Description  Login with google
// @Param 		 request body request.AssignRoleUser true "login param"
// @Param 		 id path string  true  "User ID"
// @Tags         User API
// @Produce      json
// @Success		 200	{object} bool
// @Failure		 400	{object} error
// @Router       /:id/assign-role [post]
func (s Service) LoginByGoogle(ctx context.Context, email string) (*entity.User, error) {
	// generate url call to google

	//
	panic("implement me")
}
