package httpapi

import (
	"net/http"
	"sendo/internal/category/service"
	"sendo/internal/category/service/request"
	"sendo/pkg/constants"
	"sendo/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

type apiController struct {
	service service.CategoryUseCase
}

func NewAPIController(s service.CategoryUseCase) apiController {
	return apiController{service: s}
}

func (api apiController) AddCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		var categoryRequest request.CategoryCreateRequest

		err := c.ShouldBindJSON(&categoryRequest)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Param post fail",
			})
			return
		}

		result, errCreate := api.service.AddCategory(c, categoryRequest)

		if errCreate != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Create category fail",
			})
			return
		}
		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, "Create success"))
	}
}

func (api apiController) GetListCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		var queryParam request.QueryParam

		err := c.ShouldBindQuery(&queryParam)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": constants.ParamPassedWrong,
			})
			return
		}

		result, errGetList := api.service.GetListCategory(c, queryParam)
		if errGetList != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": constants.WrongGetListCategory,
			})
			return
		}

		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, constants.GetListCategoryDone))
	}
}

func (api apiController) GetCategoryById() func(c *gin.Context) {
	return func(c *gin.Context) {

		var queryParam request.RequestURI

		err := c.ShouldBindUri(&queryParam)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Query param fail",
			})
			return
		}

		result, errGetCat := api.service.GetCategoryById(c, queryParam.ID)
		if errGetCat != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Get cat fail",
			})
			return
		}

		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, "Get cat done"))
	}
}

func (api apiController) SearchCategory() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func (api apiController) GetListAndChild() func(c *gin.Context) {
	return func(c *gin.Context) {
		var queryParam request.QueryParam

		err := c.ShouldBindQuery(&queryParam)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Query param fail",
			})
			return
		}
		result, errGetList := api.service.GetCategoryAndChild(c, queryParam)
		if errGetList != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Query param fail",
			})
			return
		}

		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, "Get list done"))
	}
}
