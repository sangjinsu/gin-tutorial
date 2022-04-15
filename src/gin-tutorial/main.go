package main

import (
	"gin-tutorial/controller"
	"gin-tutorial/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.Use(gin.Logger())

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})

	server.Run(":8080")
}
