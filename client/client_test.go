package client

import (
	"fmt"
	"testing"
)

func TestGetURL_GetHttp(t *testing.T) {
	g := GetURL{}

	type Getinfo struct {
		Province string
		Age      int
	}

	type GetinfoList struct {
		List []Getinfo
	}

	ClientInfo := GetinfoList{[]Getinfo{
		{Province: "Test1", Age: 1},
		{Province: "Test2", Age: 2},
		{Province: "Test1", Age: 3},
	},
	}
	fmt.Println(ClientInfo)
	PvS := g.ProvinceService()

	// AgS := g.AgeService()

	if PvS != nil {
		t.Errorf("got %v ,want %v", PvS, nil)
	}

}
