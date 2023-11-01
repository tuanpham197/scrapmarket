package middleware

import (
	"github.com/gin-gonic/gin"
)

func RoleAdminMiddleware(c *gin.Context) {
	// userId := c.GetString("userId")

	c.Next()
}
