package main

import(
	//"./controllar"
	//"github.com/gin-gonic/gin"
	"log"
)

func main(){

	if err := Router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}