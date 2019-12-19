package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	//事前にテンプテートの実行HTMLを指定
	router.LoadHTMLGlob("../views/*.html")
	router.GET("/", Home)
	router.GET("/Send", Send)
	router.Run(":8080")
}

func Home(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "base.html", gin.H{})
}

//Send は
func Send(c *gin.Context) {
	// mail := c.Param("mail")
	// pass := c.Param("pass")
	c.JSON(http.StatusOK, gin.H{
		"mail": "mail",
		// "pass": pass,
	})
}
