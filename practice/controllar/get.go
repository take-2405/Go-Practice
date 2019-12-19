package controllar

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(cxt *gin.Context) {
	cxt.HTML(http.StatusOK, "base.html", gin.H{})
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
