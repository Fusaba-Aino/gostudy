package github_com_gin_gonic_gin

import "github.com/gin-gonic/gin"

func Controller() {
	//注意，一般不会这么写，会写handler

	r := gin.Default()
	//get测试
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get_ping",
		})
	})

	//POST测试
	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post_ping",
		})
	})

	//RESTFUL GET传参
	r.GET("/api/in/:in", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Param("in"),
		})
	})

	//RESTFUL POST传参,但我觉得一般不会这么用
	r.POST("/api/in/:in", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Param("in"),
		})
	})

	//POST传参
	r.POST("/api/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.PostForm("test"),
		})
	})

	//GET传参
	r.GET("/api/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Query("test"),
		})
	})

	admin := r.Group("/admin")
	admin.Use(Auth())
	//测试中间件
	admin.POST("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "admin",
		})
	})

	err := r.Run(":8089")
	if err != nil {
		return
	}
}
