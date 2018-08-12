package main

import (
	"testing"
)

func Test_Add(t *testing.T) {
	type check struct {
		p     permission
		valid bool
	}
	tests := []struct {
		name    string
		nm      string
		include permission
		exclude permission
		want    []check
	}{

		{"1", "prabesh", permission{Country: "IN", State: "KL"}, permission{Country: "IN", State: "KL", Province: "VRKLA"},
			[]check{
				check{permission{"IN", "KL", ""}, false},
				check{permission{"IN", "KL", "DEVUA"}, true},
				check{permission{"IN", "TN", "KLRAI"}, false},
			}},
		{"2", "prabesh", permission{Country: "IN", State: "KL", Province: "VRKLA"}, permission{},
			[]check{
				check{permission{"IN", "KL", "VRKLA"}, true},
				check{permission{"IN", "KL", "DEVUA"}, true},
				check{permission{"IN", "TN", "KLRAI"}, false},
			}},
		{"3", "prabesh", permission{Country: "IN"}, permission{Country: "IN", State: "KL"},
			[]check{
				check{permission{"IN", "KL", "VRKLA"}, false},
				check{permission{"IN", "KL", "TPPPPP"}, false},
				check{permission{"IN", "TN", "KLRAI"}, true},
			}},
		{"4", "prajesh", permission{Country: "IN"}, permission{Country: "IN", State: "KL"},
			[]check{
				check{permission{"IN", "KL", "VRKLA"}, false},
				check{permission{"IN", "KL", "TPPPPP"}, false},
				check{permission{"IN", "TN", "KLRAI"}, true},
			}},
		{"5", "prajesh", permission{Country: "IN", State: "KL"}, permission{Country: "IN", State: "KL", Province: "VRKLA"},
			[]check{
				check{permission{"IN", "KL", "VRKLA"}, false},
				check{permission{"IN", "KL", "DEVUA"}, true},
				check{permission{"IN", "TN", "KLRAI"}, true},
			},
		},
		//testing for add function
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			distributor := getDist(tt.nm)
			distributor.Add(tt.include, tt.exclude)
			for _, r := range tt.want {
				if got := distributor.Check(r.p); got != r.valid {
					t.Errorf("check() = %v,for permssion %s want  %v", got, r.p, r.valid)
				}
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
