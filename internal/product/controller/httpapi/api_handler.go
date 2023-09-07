package httpapi

import (
	"fmt"
	"net/http"
	authMiddleware "sendo/internal/auth/controller/middleware"
	"sendo/internal/product/service"
	"sendo/internal/product/service/request"
	shopMiddleware "sendo/internal/shop/controller/middleware"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService service.ProductUseCase
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
		fmt.Println("FILTER :::", catParam)

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

func (api productController) SetUpRoute(group *gin.RouterGroup) {

	group.GET("/", api.GetListProduct())

	group.Use(authMiddleware.TokenVerificationMiddleware)
	group.Use(shopMiddleware.ShopMiddleware)
	group.POST("/create", api.InsertProduct())
}
