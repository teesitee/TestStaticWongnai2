package main

import (
	"fmt"
	"teststaticwongnai2/api"
	"teststaticwongnai2/client"
	"teststaticwongnai2/router"
)

func main() {
	fmt.Println("Hello")

	var svc client.GetURL

	h := api.Handler{
		Serv: svc,
	}
	r := router.Newrouter(h)
	r.Run(":8080")
}
