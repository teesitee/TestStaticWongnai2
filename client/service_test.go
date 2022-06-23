package client

import (
	"reflect"
	"testing"
)

func TestGetURL_AgeService(t *testing.T) {
	tests := []struct {
		name string
		g    GetURL
		want Age
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.AgeService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetURL.AgeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetURL_ProvinceService(t *testing.T) {
	type args struct {
		get GetinfoList
	}
	tests := []struct {
		name string
		g    GetURL
		args args
		want map[string]int
	}{
		{"test", GetURL{}, args{}, map[string]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.ProvinceService(tt.args.get); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetURL.ProvinceService() = %v, want %v", got, tt.want)
			}
		})
	}
}
