package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Сообщение-": "вам!"})
	})

	r.Run(":30003")
}
