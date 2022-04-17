package main

import (
	"gin-tutorial/controller"
	"gin-tutorial/middleware"
	"gin-tutorial/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

//func setupLogOutput() {
//	f, _ := os.Create("request.log")
//	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
//}

func main() {

	//setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})

	server.Run(":8080")
}
