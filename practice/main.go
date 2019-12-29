package main

import(
	//"./controllar"
	//"github.com/gin-gonic/gin"
	"./router"
	"log"
)

func main(){
router.Init()
	if err := router.Router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}