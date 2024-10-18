package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCSO() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isCSO, exists := c.Get("isCSO"); !exists || !isCSO.(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
