//Enes Can Güneş - kodlib.com - 2021
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test gin framework"})
	})
	router.Run(":3000")
}
