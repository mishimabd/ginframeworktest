package main

import (
	"github.com/gin-gonic/gin"
	"test.go/controller"
	"test.go/service"
)

var VideoService service.VideoService = service.New()
var VideoController controller.VideoController = controller.New(VideoService)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
