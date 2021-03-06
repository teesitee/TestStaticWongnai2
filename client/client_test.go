package client

import (
	"fmt"
	"testing"
)

var P map[string]int

func TestGetURL_GetHttp(t *testing.T) {
	g := GetURL{}

	testClientInfo := GetinfoList{[]Getinfo{
		{Province: "Test1", Age: 1},
		{Province: "Test2", Age: 2},
		{Province: "Test1", Age: 3},
	},
	}

	P = g.ProvinceService(testClientInfo)

	if len(P) == 0 {
		t.Errorf("got %v ", P)
	} else {
		fmt.Println("Print test :", testClientInfo)
	}

}

var A Age

func TestGetURL_GetHttpAge(t *testing.T) {
	g := GetURL{}

	testClientInfo := GetinfoList{[]Getinfo{
		{Province: "Test1", Age: 1},
		{Province: "Test2", Age: 2},
		{Province: "Test1", Age: 1},
	},
	}
	A = g.AgeService(testClientInfo)

	if A.Age0to30 != 3 {
		t.Errorf("got %v ", A)
	} else if A.Age31to60 != 0 {
		t.Errorf("got %v ", A)
	} else if A.Age61up != 0 {
		t.Errorf("got %v ", A)
	} else if A.AgeNA != 0 {
		t.Errorf("got %v ", A)
	}
}
