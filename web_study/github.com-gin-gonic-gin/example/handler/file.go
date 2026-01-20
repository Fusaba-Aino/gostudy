package handler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"msg": err.Error()})
		return
	}
	if f.Size == 0 {
		c.AbortWithStatusJSON(400, gin.H{"msg": "empty file"})
		return
	}

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"msg": err.Error()})
		return
	}

	filename := filepath.Base(f.Filename) // 防止路径穿越
	dstPath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(f, dstPath); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "ok", "path": dstPath, "size": f.Size})
}

func FileDownload(c *gin.Context) {
	filename := filepath.Base(c.Param("filename"))
	filePath := filepath.Join("./uploads", filename)
	fmt.Println("filePath:", filePath)
	c.File(filePath)
}
