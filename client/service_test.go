package client

import (
	"reflect"
	"testing"
)

func TestGetURL_ProvinceService(t *testing.T) {
	tests := []struct {
		name string
		g    GetURL
		want map[string]int
	}{
		{"TestProvinceService", GetURL{}, map[string]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.ProvinceService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetURL.ProvinceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetURL_AgeService(t *testing.T) {
	tests := []struct {
		name string
		g    GetURL
		want Age
	}{
		{"TestAgeService", GetURL{}, Age{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.AgeService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetURL.AgeService() = %v, want %v", got, tt.want)
			}
		})
	}
}
