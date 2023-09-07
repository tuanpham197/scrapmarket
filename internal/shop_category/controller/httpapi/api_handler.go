package httpapi

import (
	"fmt"
	"net/http"
	authMiddleware "sendo/internal/auth/controller/middleware"
	shopMiddleware "sendo/internal/shop/controller/middleware"
	"sendo/internal/shop_category/service"
	"sendo/internal/shop_category/service/request"
	"sendo/pkg/queue"
	"sendo/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

type apiController struct {
	shopCategoryService service.ShopCategoryUseCase
}

func NewAPIController(s service.ShopCategoryUseCase) apiController {
	return apiController{shopCategoryService: s}
}

func (api apiController) CreateShopCategory() func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var createShopCategory request.CreateShopCategory

		errBind := c.ShouldBindJSON(&createShopCategory)
		if errBind != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": errBind.Error(),
			})
			return
		}

		shopId, _ := c.Get("shopId")

		createShopCategory.ShopId, _ = shopId.(string)

		result, err := api.shopCategoryService.CreateShopCategory(c, createShopCategory)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func (api apiController) UpdateShopCategory() func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var updateShopCategory request.UpdateShopCategory

		errBind := c.ShouldBindJSON(&updateShopCategory)
		if errBind != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": errBind.Error(),
			})
			return
		}

		shopId := c.GetString("shopId")
		categoryId := c.Param("id")

		result, err := api.shopCategoryService.UpdateShopCategory(c, updateShopCategory, shopId, categoryId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func (api apiController) DeleteShopCategory() func(ctx *gin.Context) {
	return func(c *gin.Context) {

		shopId := c.GetString("shopId")
		categoryId := c.Param("id")

		result, err := api.shopCategoryService.DeleteShopCategory(c, shopId, categoryId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func (api apiController) GetListShopCategoryAndProduct() func(ctx *gin.Context) {
	return func(c *gin.Context) {

		shopId := c.Param("id")
		result, err := api.shopCategoryService.GetListShopCategoryAndProduct(c, shopId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func (api apiController) TestQueue() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("123")
		ch := queue.RabbitChannel

		q, err := ch.QueueDeclare(
			"hello-queue", // name
			false,         // durable
			false,         // delete when unused
			false,         // exclusive
			false,         // no-wait
			nil,           // arguments
		)

		queue.Publish(ch, q.Name, "Test messsage")

		response.QueueFailOnError(err, "Failed to publish a message")

		c.JSON(http.StatusOK, gin.H{
			"message": "Task-1 Received Successfully",
		})
	}
}

func (api apiController) SetUpRoute(group *gin.RouterGroup) {
	group.GET("/test-queue", api.TestQueue())
	group.GET("/:id", api.GetListShopCategoryAndProduct())
	group.Use(authMiddleware.TokenVerificationMiddleware)
	group.Use(shopMiddleware.ShopMiddleware)
	group.POST("/", api.CreateShopCategory())
	group.PATCH("/:id", api.UpdateShopCategory())
	group.DELETE("/:id", api.DeleteShopCategory())
}
