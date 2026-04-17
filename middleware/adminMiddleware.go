package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		user, _ := c.Get("user")
		claims := user.(map[string]interface{})

		if claims["role"] != "admin" { // SCM secured role based security //
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin only"})
			c.Abort()
			return
		}

		c.Next()
	}
}