package github_com_gin_gonic_gin

import (
	"github.com/gin-gonic/gin"
)

// 中间件类似于java的filtter,这是个简单的样例
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, _ := c.Cookie("test")
		if cookie == "123456" {
			c.Next()
		}
		c.Abort()
	}
}
