package httpapi

import (
	"net/http"
	"sendo/internal/auth/controller/middleware"
	"sendo/internal/auth/service"
	"sendo/internal/auth/service/request"
	"sendo/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

type apiController struct {
	service service.UserUseCase
}

func NewAPIController(s service.UserUseCase) apiController {
	return apiController{service: s}
}

func (api apiController) Login() func(ctx *gin.Context) {

	return func(c *gin.Context) {

		var userLogin request.UserLogin
		errBind := c.ShouldBindJSON(&userLogin)

		if errBind != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": errBind.Error(),
			})
			return
		}
		result, err := api.service.Login(c, userLogin)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, ""))
	}
}

func (api apiController) Register() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var userRegister request.UserRegister

		errBind := c.ShouldBindJSON(&userRegister)
		if errBind != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": errBind,
			})
			return
		}
		result, err := api.service.Register(c, userRegister)
		if err != nil {
			if _, ok := err.(request.UserExistsError); ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, ""))
	}
}

func (api apiController) GetInfoUser() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		userId := c.GetString("userId")

		user, err := api.service.GetInfoUser(c, userId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response.ResponseData(user, http.StatusOK, ""))
	}

}

func (api apiController) RefreshToken() func(c *gin.Context) {
	return func(c *gin.Context) {
		token, _ := c.GetPostForm("refresh_token")

		tokenRequest := request.RefreshTokenRequest{
			RefreshToken: token,
		}

		result, err := api.service.RefreshToken(c, tokenRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, ""))
	}
}

func (api apiController) AssignRoleUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		// check role current user
		userId := c.GetString("userId")
		user, errUser := api.service.GetInfoUser(c, userId)
		if errUser != nil {
			c.JSON(http.StatusOK, gin.H{
				"data":    errUser.Error(),
				"message": "",
			})
			return
		}
		if !user.HasRole("admin") && !user.HasPermission("assign role") {
			c.JSON(http.StatusOK, gin.H{
				"data":    "user not have permission",
				"message": "",
			})
			return
		}
		// end check
		var userRequest request.UserRoleRequest

		if err := c.ShouldBindUri(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		var request request.AssignRoleUser
		errBind := c.ShouldBindJSON(&request)
		if errBind != nil {
			c.JSON(http.StatusOK, gin.H{
				"data":    errBind.Error(),
				"message": "",
			})
			return
		}
		if len(request.Roles) < 1 {
			c.JSON(http.StatusOK, gin.H{
				"data":    "please pass slice not null",
				"message": "",
			})
			return
		}

		result, err := api.service.AssignRoleUser(c, request, userRequest.ID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, ""))
	}
}

func (api apiController) LoginGoogle() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, response.ResponseData([]int{}, http.StatusOK, ""))
	}
}

func (api apiController) CalllBackLoginGoogle() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, response.ResponseData([]int{}, http.StatusOK, ""))
	}
}

func (api apiController) SetUpRoute(group *gin.RouterGroup) {
	// middleware
	// group.Use();
	group.POST("/login", api.Login())
	group.POST("/login-google", api.LoginGoogle())
	group.POST("/login-google/callback", api.CalllBackLoginGoogle())
	group.POST("/register", api.Register())
	group.POST("/refresh-token", api.RefreshToken())

	group.Use(middleware.TokenVerificationMiddleware)
	group.GET("/users/info", api.GetInfoUser())
	group.POST("/:id/assign-role", api.AssignRoleUser())
}
