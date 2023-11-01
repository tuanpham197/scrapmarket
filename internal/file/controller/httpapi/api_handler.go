package httpapi

import (
	"net/http"
	"sendo/internal/file/service"
	"sendo/internal/file/service/requests"
	"sendo/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

type apiController struct {
	service service.FileUseCase
}

func NewAPIController(s service.FileUseCase) apiController {
	return apiController{service: s}
}

func (api apiController) UploadImage() func(c *gin.Context) {

	return func(c *gin.Context) {
		var fileForm requests.FileImage
		err := c.ShouldBind(&fileForm)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		result, err := api.service.UploadImage(c.Request.Context(), fileForm)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, ""))
	}
}

func (api apiController) RemoveImage() func(c *gin.Context) {

	return func(c *gin.Context) {
		var req requests.RequestDelete
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		result, err := api.service.RemoveImage(c.Request.Context(), req.Path)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, ""))
	}
}

func (api apiController) UpdateImageMultiple() func(c *gin.Context) {
	return func(c *gin.Context) {
		var req requests.FileImageMultiple
		errBind := c.ShouldBind(&req)

		if errBind != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errBind.Error()})
			return
		}

		result, err := api.service.UploadMultipleImage(c, req)

		if err != nil {
			if errValidate, ok := err.(*requests.ValidateImageError); ok {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"data":    errValidate.Error(),
					"message": "",
					"code":    http.StatusUnprocessableEntity,
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"data":    err.Error(),
				"message": "",
				"code":    http.StatusBadRequest,
			})
			return
		}

		c.JSON(http.StatusOK, response.ResponseData(result, http.StatusOK, ""))
	}
}

func (api apiController) SetUpRoute(group *gin.RouterGroup) {
	// middleware
	// group.Use();
	group.POST("/image", api.UploadImage())
	group.DELETE("/image", api.RemoveImage())
	group.POST("/upload-multiple", api.UpdateImageMultiple())
}
