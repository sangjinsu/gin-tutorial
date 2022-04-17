package controller

import (
	"gin-tutorial/entity"
	"gin-tutorial/service"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"strings"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func New(service service.VideoService) VideoController {
	validate = validator.New(&validator.Config{
		TagName:      "validate",
		FieldNameTag: "title",
	})
	validate.RegisterValidation("is-cool", func(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
		return strings.Contains(field.String(), "is-cool")
	})
	return &controller{service: service}
}
