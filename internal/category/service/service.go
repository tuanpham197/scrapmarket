package service

import (
	"context"
	"sendo/internal/category/service/entity"
	"sendo/internal/category/service/request"
	"sendo/pkg/utils/paginations"
)

type service struct {
	repository CategoryRepository
}

func NewService(repository CategoryRepository) service {
	return service{repository: repository}
}

// Add category
// @Summary      Add category
// @Description  Add category
// @Param 		 request body request.CategoryCreateRequest true "create param"
// @Tags         Category
// @Produce      json
// @Success		 200	{object} entity.Category
// @Failure		 400	{object} error
// @Router       /categories [post]
func (s service) AddCategory(ctx context.Context, request request.CategoryCreateRequest) (*entity.Category, error) {
	result, err := s.repository.Insert(ctx, request.Name, request.Thumbnail, request.ParentId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get list category
// @Summary      Get list category
// @Description  Get list category
// @Param 		 request body request.QueryParam true "get list param"
// @Tags         Category
// @Produce      json
// @Success		 200	{object} paginations.Pagination
// @Failure		 400	{object} error
// @Router       /categories [get]
func (s service) GetListCategory(ctx context.Context, query request.QueryParam) (*paginations.Pagination, error) {
	result, err := s.repository.GetList(ctx, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get detail category by id
// @Summary      Get detail category by id
// @Description  Get detail category by id
// @Param 		 id path string true "Category id"
// @Tags         Category
// @Produce      json
// @Success		 200	{object} entity.Category
// @Failure		 400	{object} error
// @Router       /categories/:id [get]
func (s service) GetCategoryById(ctx context.Context, id string) (*entity.Category, error) {
	result, err := s.repository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s service) SearchCategory(ctx context.Context, keyword string) ([]*entity.Category, error) {
	return nil, nil
}
