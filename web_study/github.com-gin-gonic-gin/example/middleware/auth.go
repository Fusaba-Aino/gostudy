package middleware

import "github.com/gin-gonic/gin"

func AuthTest(c *gin.Context) {
	auth, _ := c.Cookie("auth")
	if auth == "admin" {
		c.Next()
		return
	}
	c.JSON(403, gin.H{
		"code": 403,
	})
	c.Abort()
}
