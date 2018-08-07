package main

import (
	"testing"
)

func Test_Check(t *testing.T) {
	DistData["prabesh"] = Dist{Name: "prabesh", Country: map[string]Country{
		"IN": Country{
			State: map[string]State{
				"IN:all": State{},
				"KL": State{
					Province: map[string]Province{
						"VRKLA": Province{},
					},
				},
				"TN": State{
					Province: map[string]Province{
						"WLGTN": Province{},
					},
				},
			},
		},
	}}
	tests := []struct {
		name    string
		nm      string
		include string
		want    bool
	}{
		{"simple", "prabesh", "IN,KL,VRKLA", false},
		{"Success", "prabesh", "IN,KL,DEVUA", true},
		{"Faliure case", "prabesh", "IN,TN,KLRAI", true},
		{"Invalid distributor", "prabeshf", "IN,TN,WLGTN", false},
	}
	for _, tt := range tests {
		dis := DistData[tt.nm]
		t.Run(tt.name, func(t *testing.T) {
			if got := dis.Check(tt.include); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Add(t *testing.T) {

	tests := []struct {
		name    string
		nm      string
		include string
		exclude string
		check   string
		want    bool
	}{

		{"Adding permission to India except kerala", "prabesh", "IN", "KL", "IN,KL,DEVUA", false},
		{"Adding permission to India-Kerala except VRKLA", "prabesh", "IN,KL", "VRKLA", "IN,KL,DEVUA", false},
		{"Adding permission to India-kerala-DEVUA", "prabesh", "IN,KL,DEVUA", "", "IN,KL,DEVUA", true},

		//testing for add function
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getDist(tt.nm).Add(tt.include, tt.exclude, tt.nm)
			if got := getDist(tt.nm).Check(tt.check); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_AddSub(t *testing.T) {
	tests := []struct {
		name    string
		nm      string
		include string
		exclude string
		check   string
		want    bool
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
