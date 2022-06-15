package client

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type GetURL struct {
}

type Getinfo struct {
	Province string `json:"Province"`
	Age      int    `json:"Age"`
}

type GetinfoList struct {
	List []Getinfo `json:"Data"`
}

var ClientInfo GetinfoList

func (g GetURL) GetHttp() (map[string]int, Age) {

	resp, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		log.Fatalf("Cannot GetHttp %s", err)
	}
	bodybyte, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatalf("Cannot ReadAll %s", err2)
	}
	err3 := json.Unmarshal(bodybyte, &ClientInfo)
	if err3 != nil {
		log.Fatalf("Cannot Unmarshal %s", err3)
	}

	return g.ProvinceService(), g.AgeService()
}
