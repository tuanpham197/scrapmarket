package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"sendo/internal/auth/service/entity"
	"sendo/internal/auth/service/request"
	"sendo/pkg/common"
	"sendo/pkg/constants"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	dataUser = []entity.User{
		{
			Id:        uuid.MustParse("20e4b3e4-6e3e-460a-b71d-5755ebefa317"),
			UserName:  "user_one",
			LastName:  "last_one",
			Email:     "admin@gmail.com",
			FirstName: "first_one",
			Password:  "$2a$10$12wyaE4R9jhbvT.KFQeN4uvCoiyjH6bw408oNtXqvtVDYI.HyhMMO",
			Salt:      "Z-80b",
		},
		{
			Id:        uuid.MustParse("20e4b3e4-6e3e-460a-b71d-5755ebefa318"),
			UserName:  "user_two",
			LastName:  "last_two",
			Email:     "admin2@gmail.com",
			FirstName: "two",
			Password:  "$2a$10$12wyaE4R9jhbvT.KFQeN4uvCoiyjH6bw408oNtXqvtVDYI.HyhMMO",
			Salt:      "Z-80b",
		},
	}

	userLogin = request.UserLogin{
		Email:    "admin@gmail.com",
		Password: "12345",
	}

	userLoginNotExist = request.UserLogin{
		Email:    "admin223@gmail.com",
		Password: "123456",
	}

	userLoginWrongPass = request.UserLogin{
		Email:    "admin@gmail.com",
		Password: "123456",
	}

	userRegister = request.UserRegister{
		UserName: "test@gmail.com",
		Password: "12345",
	}

	userRegisterExists = request.UserRegister{
		UserName: "admin2@gmail.com",
		Password: "12345",
	}
)

type mockRepository struct{}

func (m mockRepository) Insert(ctx context.Context, userInfo request.UserRegister) (*entity.User, error) {
	user := entity.User{
		Email:    userInfo.Email,
		Password: userInfo.Password,
	}
	dataUser = append(dataUser, user)
	return &user, nil
}

func (m mockRepository) GetOne(ctx context.Context, id string) (*entity.User, error) {
	for _, user := range dataUser {
		if user.Id.String() == id {
			return &user, nil
		}
	}
	return nil, constants.ErrUserNotExist
}

func (m mockRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockRepository) Delete(ctx context.Context, id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	for _, user := range dataUser {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, constants.ErrUserNotExist
}

func (m mockRepository) AssignRoleUser(ctx context.Context, roleId request.AssignRoleUser, userId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func Test_Service_Login(t *testing.T) {

	errEnv := godotenv.Load("/home/nlcpu0203/source/me/sendo/scrapmarketbe/.env")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	repo := mockRepository{}
	hasher := new(common.Hasher)
	logger := common.SugarLog()

	authSer := NewService(repo, hasher, logger)

	Convey("Given some data to login", t, func() {
		testTable := []struct {
			userLogin   request.UserLogin
			expectError error
			expectData  interface{}
		}{
			{
				userLogin:   userLogin,
				expectError: nil,
				expectData: request.UserLoginResponse{
					AccessToken:  "",
					RefreshToken: "",
					UserInfo: request.UserInfo{
						Email: userLogin.Email,
					},
				},
			},
			{
				userLogin:   userLoginWrongPass,
				expectError: constants.ErrWrongPass,
				expectData:  nil,
			},
			{
				userLogin:   userLoginNotExist,
				expectError: constants.ErrUserNotExist,
				expectData:  nil,
			},
		}

		Convey("Start login test", func() {
			for i, item := range testTable {

				data, err := authSer.Login(context.Background(), item.userLogin)
				conveySuit := fmt.Sprintf("compare with case %d", i)

				Convey(conveySuit, func() {
					if item.expectData != nil {
						switch item.expectData.(type) {
						case request.UserLoginResponse:
							So(err, ShouldEqual, nil)
							So(data.AccessToken, ShouldNotBeEmpty)
							So(data.RefreshToken, ShouldNotBeEmpty)
							So(data.UserInfo, ShouldNotBeEmpty)
						}
					}
					if errors.Is(item.expectError, constants.ErrUserNotExist) {
						So(err.Error(), ShouldEqual, constants.ErrUserNotExist.Error())
					}

					if errors.Is(item.expectError, constants.ErrWrongPass) {
						So(err.Error(), ShouldEqual, constants.ErrWrongPass.Error())
					}
				})
			}
		})
	})

}

func TestService_Register(t *testing.T) {
	errEnv := godotenv.Load("/home/nlcpu0203/source/me/sendo/scrapmarketbe/.env")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	repo := mockRepository{}
	hasher := new(common.Hasher)
	logger := common.SugarLog()

	authSer := NewService(repo, hasher, logger)

	Convey("Given some data to register", t, func() {
		testTable := []struct {
			userRegister request.UserRegister
			expectError  error
			expectData   interface{}
		}{
			{
				userRegister: userRegister,
				expectError:  nil,
				expectData: request.UserLoginResponse{
					AccessToken:  "",
					RefreshToken: "",
					UserInfo: request.UserInfo{
						Email: userLogin.Email,
					},
				},
			},
			{
				userRegister: userRegisterExists,
				expectError:  constants.ErrWrongPass,
				expectData:   nil,
			},
		}

		Convey("Start register test case", func() {
			for i, item := range testTable {

				data, err := authSer.Register(context.Background(), item.userRegister)
				conveySuit := fmt.Sprintf("compare with case %d", i)

				Convey(conveySuit, func() {
					if item.expectData != nil {
						switch item.expectData.(type) {
						case request.UserLoginResponse:
							So(err, ShouldEqual, nil)
							So(data.AccessToken, ShouldNotBeEmpty)
							So(data.RefreshToken, ShouldNotBeEmpty)
							So(data.UserInfo, ShouldNotBeEmpty)
						}
					}
				})
			}
		})
	})
}

func TestService_AssignRoleUser(t *testing.T) {

}

func TestService_RefreshToken(t *testing.T) {
	errEnv := godotenv.Load("/home/nlcpu0203/source/me/sendo/scrapmarketbe/.env")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	repo := mockRepository{}
	hasher := new(common.Hasher)
	logger := common.SugarLog()

	authSer := NewService(repo, hasher, logger)
	Convey("Given some data to refresh token", t, func() {
		testTable := []struct {
			userLogin   request.UserLogin
			expectError error
			expectData  interface{}
		}{
			{
				userLogin:   userLogin,
				expectError: nil,
				expectData: request.UserLoginResponse{
					AccessToken:  "",
					RefreshToken: "",
					UserInfo: request.UserInfo{
						Email: userLogin.Email,
					},
				},
			},
		}

		Convey("Start Refresh token test", func() {
			for i, item := range testTable {

				data, err := authSer.Login(context.Background(), item.userLogin)
				conveySuit := fmt.Sprintf("compare with case %d", i)
				if err == nil {
					userRefresh, errRefresh := authSer.RefreshToken(context.Background(), request.RefreshTokenRequest{
						RefreshToken: data.RefreshToken,
					})

					Convey(conveySuit, func() {
						if item.expectData != nil {
							switch item.expectData.(type) {
							case request.UserLoginResponse:
								So(errRefresh, ShouldEqual, nil)
								So(userRefresh.RefreshToken, ShouldNotBeEmpty)
								So(userRefresh.AccessToken, ShouldNotBeEmpty)
								So(userRefresh.UserInfo, ShouldNotBeEmpty)
							}
						}
					})

				}

			}
		})
	})
}
