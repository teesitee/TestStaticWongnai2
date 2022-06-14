package main

import (
	"fmt"
	"teststaticwongnai2/api"
	"teststaticwongnai2/client"
	"teststaticwongnai2/router"
	"teststaticwongnai2/service"
)

func main() {
	fmt.Println("Hello")

	var rep service.ServiceRepo
	svc := client.GetURL{
		Repo: rep,
	}
	h := api.Handler{
		Serv: svc,
	}
	r := router.Newrouter(h)
	r.Run(":8080")
}
