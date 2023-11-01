package service

import (
	"context"
	"sendo/internal/product/service/request"
)

type configValueService struct {
	configValueRepository ConfigValueRepository
}

func NewConfigValueService(configValueRepository ConfigValueRepository) configValueService {
	return configValueService{configValueRepository: configValueRepository}
}

// Create      Create
// @Summary      Create config value
// @Description  Create config value
// @Param 		 request body request.CreateConfigValue true "create param"
// @Tags         config
// @Produce      json
// @Success		 200	{object} request.CreateConfigValueResponse
// @Failure		 400	{object} error
// @Router       /config-values/create [post]
func (c configValueService) CreateConfigValue(ctx context.Context, createConfigValue request.CreateConfigValue) (*request.CreateConfigValueResponse, error) {

	// Insert config value
	result, err := c.configValueRepository.InsertConfigValue(ctx, createConfigValue)
	if err != nil {
		return nil, err
	}

	return &request.CreateConfigValueResponse{
		Id:       result.Id.String(),
		Name:     result.Name,
		ConfigId: result.ConfigId,
	}, nil
}
