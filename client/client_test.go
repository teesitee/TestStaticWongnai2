package client

import (
	"reflect"
	"testing"
)

func TestGetURL_GetHttp(t *testing.T) {

	tests := []struct {
		name  string
		g     GetURL
		want  map[string]int
		want1 Age
	}{
		{"TestGetHttp", GetURL{}, PvS, AgS},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.g.GetHttp()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetURL.GetHttp() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetURL.GetHttp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
