package mysql

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"sendo/internal/auth/service"
	"sendo/internal/auth/service/entity"
	"sendo/internal/auth/service/request"
	permissionEntity "sendo/internal/permission/service/entity"
)

type mysqlRepo struct {
	db   *gorm.DB
	reds *redis.Client
}

func NewMySQLRepo(db *gorm.DB, reds *redis.Client) service.UserRepository {
	return mysqlRepo{db: db, reds: reds}
}

func (repo mysqlRepo) GetOne(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := repo.db.Preload("Roles").Preload("Permissions").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo mysqlRepo) GetAll(ctx context.Context) ([]entity.User, error) {
	panic("implement me")
}
func (repo mysqlRepo) Insert(ctx context.Context, userRegister request.UserRegister) (*entity.User, error) {

	user := entity.User{
		UserName:  userRegister.UserName,
		LastName:  userRegister.LastName,
		FirstName: userRegister.FirstName,
		Email:     userRegister.Email,
		Password:  userRegister.Password,
		Salt:      userRegister.Salt,
	}

	result := repo.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (repo mysqlRepo) Delete(ctx context.Context, id string) (bool, error) {
	panic("implement me")
}

func (repo mysqlRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := repo.db.Preload("Shop").Preload("Roles").First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo mysqlRepo) AssignRoleUser(ctx context.Context, reqRole request.AssignRoleUser, userId string) (bool, error) {
	db := repo.db
	user, errUser := repo.GetOne(ctx, userId)
	if errUser != nil || user == nil {
		return false, errUser
	}

	var roles []permissionEntity.Role
	errRole := db.Find(&roles, reqRole.Roles).Error
	if errRole != nil || len(roles) < 1 {
		return false, errors.New("not found role to assign")
	}

	errAppend := db.Model(user).Omit("Roles.*").Association("Roles").Replace(&roles)
	if errAppend != nil {
		return false, errAppend
	}

	return true, nil
}
