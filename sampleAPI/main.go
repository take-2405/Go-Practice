package main

import (
	"github.com/git/Go-Practice/sampleAPI/db"
	"github.com/git/Go-Practice/sampleAPI/server"
)

func main() {
	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "HelloWorld")
	// })
	db.Init()
	//router.Run()
	db.Close()
	server.Init()
}
