package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

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
