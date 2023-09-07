package common

import (
	"net/http"
	"sendo/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Internal Server Error", err))
			}
		}()

		c.Next()
	}
}
