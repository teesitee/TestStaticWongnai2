package api

import (
	"net/http"
	"teststaticwongnai2/client"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Serv client.GetURL
}

func (h Handler) GetRespone(c *gin.Context) {
	ProviceList, AgeGroup := h.Serv.GetHttp()
	c.JSON(http.StatusOK, gin.H{"Province": ProviceList, "AgeGroup": AgeGroup})
}
