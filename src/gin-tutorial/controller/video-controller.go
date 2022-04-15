package controller

import (
	"gin-tutorial/entity"
	"gin-tutorial/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		// log.Fatal("video bind json error")
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
	c.service.Save(video)
	return video
}

func New(service service.VideoService) VideoController {
	return &controller{service: service}
}
