package httpapi

import (
	"net/http"
	authMiddleware "sendo/internal/auth/controller/middleware"
	shopMiddleware "sendo/internal/shop/controller/middleware"
	"sendo/internal/shop/service"
	"sendo/internal/shop/service/request"

	"github.com/gin-gonic/gin"
)

type apiController struct {
	shopService service.ShopUseCase
}

func NewAPIController(s service.ShopUseCase) apiController {
	return apiController{shopService: s}
}

func (api apiController) Register() func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var shopRegister request.ShopRegister

		errBind := c.ShouldBindJSON(&shopRegister)
		if errBind != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": errBind,
			})
			return
		}

		userId, _ := c.Get("userId")

		shopRegister.UserId, _ = userId.(string)

		result, err := api.shopService.ShopRegister(c, shopRegister)
		if err != nil {
			if _, ok := err.(request.ShopExistsError); ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
				return
			}
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

func (api apiController) GetShopInfo() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		userId, _ := c.Get("userId")
		shop, err := api.shopService.GetShopInfo(c, userId.(string))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": shop,
		})
	}

}

func (api apiController) SetUpRoute(group *gin.RouterGroup) {
	group.Use(authMiddleware.TokenVerificationMiddleware)
	group.POST("/register", api.Register())
	group.Use(shopMiddleware.ShopMiddleware)
	group.GET("/info", api.GetShopInfo())
}
