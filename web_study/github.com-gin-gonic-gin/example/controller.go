package example

import (
	"github.com/gin-gonic/gin"
	"gostudy/web_study/github.com-gin-gonic-gin/example/handler"
	"gostudy/web_study/github.com-gin-gonic-gin/example/middleware"
)

func TestController() {
	r := gin.Default()
	r.POST("/file", handler.FileUpload)
	r.GET("/download/:filename", handler.FileDownload)

	auth := r.Group("/admin")
	auth.Use(middleware.AuthTest)
	auth.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"username": "admin",
			},
		})
	})
	r.Run(":8080")
}
