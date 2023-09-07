package request

type CreateConfigValue struct {
	Name     string `json:"name" binding:"required"`
	ConfigId string `json:"config_id" binding:"required"`
}

type CreateConfigValueResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ConfigId string `json:"config_id"`
}
