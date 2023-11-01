package service

import (
	"context"
	"sendo/internal/product/service/request"
)

type configService struct {
	configRepository ConfigRepository
}

func NewConfigService(configRepository ConfigRepository) configService {
	return configService{configRepository: configRepository}
}

// Create      Create
// @Summary      Create config
// @Description  Create config
// @Param 		 request body request.CreateConfig true "create param"
// @Tags         config
// @Produce      json
// @Success		 200	{object} request.CreateConfigResponse
// @Failure		 400	{object} error
// @Router       /configs/create [post]
func (c configService) CreateConfig(ctx context.Context, createConfig request.CreateConfig) (*request.CreateConfigResponse, error) {

	// Insert config
	result, err := c.configRepository.InsertConfig(ctx, createConfig)
	if err != nil {
		return nil, err
	}

	return &request.CreateConfigResponse{
		Id:     result.Id.String(),
		Name:   result.Name,
		ShopId: result.ShopId,
	}, nil
}
