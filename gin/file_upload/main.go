//Enes Can Güneş - kodlib.com - 2021
//uploads/Go-Logo_Aqua.png (c) https://blog.golang.org/go-brand

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(422, gin.H{"error": "There is no file :("})
			return
		}
		c.SaveUploadedFile(headers, "./uploads/"+headers.Filename)
		c.JSON(200, headers)
	})
	router.Run(":3000")
}
