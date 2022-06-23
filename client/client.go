package client

import (
	"encoding/json"
	"fmt"
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

var PvS map[string]int
var AgS Age

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
	// fmt.Printf("Test Print : %+v\n", ClientInfo)

	PvS = g.ProvinceService(ClientInfo)
	AgS = g.AgeService()
	fmt.Printf("PvS : %+v \n", PvS)
	fmt.Printf("AgS : %+v \n", AgS)

	return PvS, AgS
}
