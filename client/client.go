package client

import (
	"encoding/json"
	"io"
	"net/http"
	"teststaticwongnai2/service"
)

type GetURL struct {
	Repo service.ServiceRepo
}

type Getinfo struct {
	Province string `json:"Province"`
	Age      int    `json:"Age"`
}

type GetinfoList struct {
	List []Getinfo `json:"Data"`
}

var ClientInfo GetinfoList

func (g GetURL) GetHttp() (map[string]int, service.Age) {

	resp, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		panic("Cannot GetInfo")
	}
	bodybyte, err := io.ReadAll(resp.Body)
	err2 := json.Unmarshal(bodybyte, &ClientInfo)
	if err2 != nil {
		panic("Cannot Unmarshal")
	}

	return g.Repo.ProvinceRepo(ClientInfo), g.Repo.AgeRepo(ClientInfo)
}
