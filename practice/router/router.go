package router

import(
	"github.com/gin-gonic/gin"
	"../controllar"
)

var(
	Router *gin.Engine
)

func Init(){
	Router := gin.Default()
	//事前にテンプテートの実行HTMLを指定
	Router.LoadHTMLGlob("template/base.html")
	Router.GET("/", controllar.Home)
	Router.GET("/Send", controllar.Send)
}

