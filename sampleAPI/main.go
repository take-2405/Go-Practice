package main

import (
	"./db"
	"./server"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "HelloWorld")
	})
	db.Init()
	router.Run()
	db.Close()
	server.Init()
}
