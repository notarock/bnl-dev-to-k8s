package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()

	myEnv := os.Getenv("VARIABLE")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/env", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Ma-var": myEnv,
		})
	})
	r.Run()
	// listen and serve on 0.0.0.0:8080
}
