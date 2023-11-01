package httpapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sendo/internal/product/service"
	"sendo/internal/product/service/request"
	"sendo/pkg/constants"
	"sendo/pkg/utils/response"
)

type productController struct {
	productService service.ProductUseCase
}

type RequestURI struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func NewAPIController(p service.ProductUseCase) productController {
	return productController{productService: p}
}

func (api productController) InsertProduct() func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var createProductRaw request.CreateProductRaw

		errBind := c.ShouldBindJSON(&createProductRaw)
		if errBind != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": errBind.Error(),
			})
			return
		}

		shopId := c.GetString("shopId")

		// Create Product
		result, err := api.productService.CreateProduct(c, createProductRaw, shopId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func (api productController) GetListProduct() func(ctx *gin.Context) {

	return func(c *gin.Context) {
		var catParam request.FilterRequest
		if err := c.ShouldBind(&catParam); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		result, errGet := api.productService.GetList(c, &catParam)
		if errGet != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": errGet.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func (api productController) GetDetail() func(ctx *gin.Context) {

	return func(c *gin.Context) {
		var requestUrl RequestURI

		errBinding := c.ShouldBindUri(&requestUrl)
		if errBinding != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, constants.BindingError, errBinding))
			return
		}

		result, err := api.productService.GetDetail(c, requestUrl.ID)

		if err != nil {
			c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, constants.NotFound, err))
			return
		}

		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, fmt.Sprintf(constants.GetDetailDone, "product")))
	}
}

func (api productController) GetConfig() func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var req request.FormRequest

		errBinding := c.ShouldBind(&req)
		if errBinding != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, constants.BindingError, errBinding))
			return
		}

		result, err := api.productService.GetConfigProduct(c, req.ProductId)

		if err != nil {
			c.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, constants.NotFound, err))
			return
		}

		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, fmt.Sprintf(constants.GetDetailDone, "product")))
	}
}
