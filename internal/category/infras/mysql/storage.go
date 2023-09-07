package mysql

import (
	"context"
	"errors"
	"sendo/internal/category/service"
	"sendo/internal/category/service/entity"
	"sendo/internal/category/service/request"
	"sendo/pkg/utils/paginations"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) service.CategoryRepository {
	return mysqlRepo{db: db}
}

func (repo mysqlRepo) Insert(ctx context.Context, name, thumbnail, parentId string) (*entity.Category, error) {

	category := entity.Category{
		Name:      name,
		Thumbnail: thumbnail,
		ParentID:  nil,
	}

	if parentId != "" {
		parentIdP := uuid.MustParse(parentId)
		category.ParentID = &parentIdP
	}

	result := repo.db.Create(&category)
	if result.Error != nil {
		return nil, errors.New("create category fail")
	}

	return &category, nil
}

func (repo mysqlRepo) GetList(ctx context.Context, queryParam request.QueryParam) (*paginations.Pagination, error) {
	query := repo.db.Model(entity.Category{})

	if queryParam.Name != "" {
		query.Where("name like ?", "%"+queryParam.Name+"%")
	}

	if queryParam.ParentId != "" {
		query.Where("parent_id = ?", queryParam.ParentId)
	}

	// return categories, nil
	page := queryParam.Page
	limit := queryParam.PerPage
	orderBy := []string{}

	var categories []entity.Category
	p := &paginations.Param{
		DB:      repo.db,
		Query:   query,
		Page:    page,
		Limit:   limit,
		OrderBy: orderBy,
	}
	result, err := paginations.Pagging(p, &categories)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TODO: Need update func to update all field
func (repo mysqlRepo) UpdateThumbnail(ctx context.Context, thumbnailPath string, category *entity.Category) (bool, error) {
	category.Thumbnail = thumbnailPath
	result := repo.db.Save(category)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (repo mysqlRepo) FindOne(ctx context.Context, id string) (*entity.Category, error) {
	db := repo.db.Debug()
	var category entity.Category
	result := db.First(&category, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}
