package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, roleExists := c.Get("role")
		if roleExists != true {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "401 Unauthorized",
			})
			return
		}
		if userRole != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "403 Forbidden",
			})
			return
		}
		c.Next()
	}
}
