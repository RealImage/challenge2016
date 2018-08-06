package main

import (
	"testing"
)

func TestCheck(t *testing.T) {

	tests := []struct {
		name    string
		nm      string
		dist    Dist
		include string
		want    bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.dist.Check(tt.include); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		nm      string
		include string
		exclude string
	}{
		{"simple", "prabesh", "IN,KL", "VRKLA"},
		{"simple", "prabesh", "IN,KL", "DEVUA"},
		{"simple", "prabesh", "IN,KL", "THRAP"},
		{"simple", "prabesh", "IN", "KL,THRAP"},
		{"simple", "prabesh", "IN,KL", "VRKLA"},
		{"simple", "prabesh", "IN,KL,THRAP", ""},
		{"simple", "prabesh", "IN,KL", "THRAP"},
		{"simple", "prabesh", "IN,KL,THRAP", ""},
		{"simple", "prabesh", "IN,KL", ""},
		{"simple", "prabesh", "IN", "KL"},
		{"simple", "prabesh", "IN,KL", ""},

		//testing for add function
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}

}
