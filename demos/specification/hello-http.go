package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET(
		"/", func(context *gin.Context) {
			context.Writer.WriteString("Hello world!")
		},
	)
	engine.Run("0.0.0.0:8888")
}
