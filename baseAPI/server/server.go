package server

import (
	"github.com/gin-gonic/gin"
	"github.com/git/Go-Practice/baseAPI/controller"
)

var (
	Router *gin.Engine
)

func init() {
	Router = gin.Default()
	Router.GET("/get", controller.Get)
	Router.POST("/post", controller.Post)

}
