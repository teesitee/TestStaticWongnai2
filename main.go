package main

import (
	"fmt"
	"teststaticwongnai2/api"
	"teststaticwongnai2/client"
	"teststaticwongnai2/router"
)

var I int

func main() {
	fmt.Println("Hello")
	I = 12
	fmt.Println("Print:", I)

	var svc client.GetURL

	h := api.Handler{
		Serv: svc,
	}
	r := router.Newrouter(h)
	r.Run(":8080")
}
