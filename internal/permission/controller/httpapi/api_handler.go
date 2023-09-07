package httpapi

import (
	"net/http"
	"sendo/internal/auth/controller/middleware"
	"sendo/internal/permission/service"
	"sendo/internal/permission/service/requests"
	"sendo/pkg/utils/validator"

	"github.com/gin-gonic/gin"
)

type apiController struct {
	roleService service.RoleUseCase
}

func NewAPIController(s service.RoleUseCase) apiController {
	return apiController{roleService: s}
}

func (api apiController) Create() func(c *gin.Context) {
	return func(c *gin.Context) {
		var req requests.RoleRequest
		errBind := c.ShouldBindJSON(&req)
		if errBind != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errBind.Error(),
			})
			return
		}

		err := validator.ValidateRequestData(req)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"errors": err,
			})
			return
		}

		result, errCreate := api.roleService.Create(c, &req)
		if errCreate != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errCreate.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    result,
			"message": "",
		})
	}
}

func (api apiController) GivePermissionTo() func(c *gin.Context) {
	return func(c *gin.Context) {
		roleId := c.GetInt("roleId")

		var req requests.AssignPermissionRequest
		errBind := c.ShouldBindJSON(&req)

		if errBind != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errBind.Error(),
			})
			return
		}

		err := validator.ValidateRequestData(req)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"errors": err,
			})
			return
		}

		result, errAssign := api.roleService.GivePermissionTo(c, uint(roleId), &req)
		if errAssign != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errAssign.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    result,
			"message": "",
		})

	}
}

func (api apiController) RevokePermissionTo() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func (api apiController) SetUpRoute(group *gin.RouterGroup) {
	// middleware
	group.Use(middleware.TokenVerificationMiddleware)
	group.POST("/", api.Create())
	group.POST("/:roleId/assign-permission", api.GivePermissionTo())
	group.POST("/revoke", api.RevokePermissionTo())
}
