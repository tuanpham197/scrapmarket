package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShopMiddleware(c *gin.Context) {
	shopId, _ := c.Get("shopId")

	if shopId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shopId"})
		c.Abort()
		return
	}

	// Call the next handler
	c.Next()
}
