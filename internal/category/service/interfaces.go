package service

import (
	"context"
	"sendo/internal/category/service/entity"
	"sendo/internal/category/service/request"
	"sendo/pkg/utils/paginations"
)

type CategoryUseCase interface {
	AddCategory(ctx context.Context, request request.CategoryCreateRequest) (*entity.Category, error)
	// handle pagination
	GetListCategory(ctx context.Context, query request.QueryParam) (*paginations.Pagination, error)
	GetCategoryById(ctx context.Context, id string) (*entity.Category, error)
	SearchCategory(ctx context.Context, keyword string) ([]*entity.Category, error)
	// GetListCategoryByParentId(ctx context.Context, parent_id string) ([]*entity.Category, error)
}

type CategoryRepository interface {
	Insert(ctx context.Context, name, thumbnail, parent_id string) (*entity.Category, error)
	GetList(ctx context.Context, query request.QueryParam) (*paginations.Pagination, error)
	UpdateThumbnail(ctx context.Context, thumbnail string, category *entity.Category) (bool, error)
	FindOne(ctx context.Context, id string) (*entity.Category, error)
}
