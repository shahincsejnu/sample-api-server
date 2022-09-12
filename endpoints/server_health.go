package endpoints

import "github.com/gin-gonic/gin"

func GetServerHealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Working fine!",
	})
}
