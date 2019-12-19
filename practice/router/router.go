package router

import(
	"github.com/gin-gonic/gin"
	"../controllar"
)

func Init(){
	Router := gin.Default()
	//事前にテンプテートの実行HTMLを指定
	Router.LoadHTMLGlob("../views/*.html")
	Router.GET("/", controllar.Home)
	Router.GET("/Send", controllar.Send)
}

