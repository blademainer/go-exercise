package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.Writer.WriteString("Hello world!")
	})
	r.GET("/hello.json", func(c *gin.Context) {
		//c.Writer.WriteString("Hello world!")
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
