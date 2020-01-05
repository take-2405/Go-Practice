package main

import (
	"fmt"
	"github.com/take-2405/sample/calculate"
)

func main(){
var a,b int
fmt.Println("a+b=sum")
fmt.Print("a=")
fmt.Scan(&a)
fmt.Print("b=")
fmt.Scan(&b)
fmt.Print("sum=")
fmt.Scan(&sum)
calculate.add(a,b,sum)
}