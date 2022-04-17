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

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/posts", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})

		apiRoutes.POST("/posts", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Video Input is Valid",
			})

		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", func(c *gin.Context) {
			videoController.ShowAll(c)
		})
	}

	server.Run(":8080")
}
