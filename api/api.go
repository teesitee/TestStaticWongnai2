package api

import (
	"net/http"
	"teststaticwongnai2/client"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=./api.go -destination=./mock_api/mock_api.go -package=mockapi_test
type HandlerMock interface {
	GetHttp() (map[string]int, client.Age)
}

type Handler struct {
	Serv HandlerMock
}

func (h Handler) GetRespone(c *gin.Context) {
	ProviceList, AgeGroup := h.Serv.GetHttp()
	c.JSON(http.StatusOK, gin.H{"Province": ProviceList, "AgeGroup": AgeGroup})
}
