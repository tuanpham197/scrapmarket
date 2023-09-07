package service

import (
	"context"
	"errors"
	"sendo/internal/auth/service/entity"
	"sendo/internal/auth/service/request"
	"sendo/pkg/common"
	"testing"
)

var (
	userOne = request.UserInfo{
		UserName:  "user_one",
		LastName:  "last_one",
		Email:     "user_one@gmail.com",
		FirstName: "first_one",
	}

	// emptyTranslation = request.UserRegister{}

	ErrDB = errors.New("db error")
)

type mockRepository struct{}

func (mockRepository) Insert(ctx context.Context, userRequest request.UserRegister) (*entity.User, error) {
	panic("test")
}

func (mockRepository) GetOne(ctx context.Context, id string) (*entity.User, error) {
	panic("test")
}

func (mockRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	panic("test")
}

func (mockRepository) Delete(ctx context.Context, id string) (bool, error) {
	panic("test")
}

func (mockRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	panic("test")
}

func (mockRepository) AssignRoleUser(ctx context.Context, reqRole request.AssignRoleUser, userId string) (bool, error) {
	panic("err")
}

func TestService_Login(t *testing.T) {

	testTable := []struct {
		userRegister request.UserRegister
		expectError  error
		expectData   request.UserLoginResponse
	}{
		{
			userRegister: request.UserRegister{
				UserName:  userOne.UserName,
				LastName:  userOne.LastName,
				FirstName: userOne.FirstName,
				Email:     userOne.Email,
			},
			expectError: request.ResponseMessageError{Message: "err"},
			expectData: request.UserLoginResponse{
				AccessToken:  "",
				RefreshToken: "",
				UserInfo:     userOne,
			},
		},
	}

	repo := mockRepository{}
	hasher := &common.Hasher{}
	authSer := NewService(repo, hasher)

	// Run test

	for _, item := range testTable {
		realData, realErr := authSer.Register(context.Background(), item.userRegister)
		if realData != &item.expectData {
			t.Errorf("Failed. Expected %v but received %v", item.expectData, realData)
		}

		if realErr != item.expectError {
			t.Errorf("Failed. Expected %v but received %v", item.expectData, realData)
		}
	}
}

func TestService_Register(t *testing.T) {

}

func TestService_RefreshToken(t *testing.T) {

}

func TestService_GetInfoUser(t *testing.T) {

}

func TestService_AssignRoleUser(t *testing.T) {

}
