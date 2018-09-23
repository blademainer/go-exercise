package main

import "github.com/gin-gonic/gin"

func main() {
	Router := gin.Default()

	Router.GET("/", )
}

func handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.Request
		request.ParseForm()
	}
}
