package router

import (
	"teststaticwongnai2/api"

	"github.com/gin-gonic/gin"
)

func Newrouter(h api.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/getinfo", h.GetRespone)
	return router

}
