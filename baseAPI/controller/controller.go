package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Get Getの処理を書く
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping",
	})
}

//Post Postの処理を書く
func Post(c *gin.Context) {
	c.Request.ParseForm()
	fmt.Println(c.Request.Form["id"])
}
