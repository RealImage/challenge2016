package service

import (
	"reflect"
	"testing"

	"github.com/saurabh-sde/challenge2016_saurabh/model"
	"github.com/saurabh-sde/challenge2016_saurabh/utils"
)

func TestAddDistributor(t *testing.T) {
	type args struct {
		req *utils.NewDistributorRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Distributor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				req: &utils.NewDistributorRequest{
					Name:     "DISTRIBUTOR1",
					Includes: []string{"IN"},
					Excludes: []string{"UP-IN"},
				},
			},
			want: &model.Distributor{
				Name:              "DISTRIBUTOR1",
				Includes:          []string{"IN"},
				Excludes:          []string{"UP-IN"},
				ParentDistributor: "",
			},
		},
		{
			name: "Test2",
			args: args{
				req: &utils.NewDistributorRequest{
					Name:     "DISTRIBUTOR2",
					Includes: []string{"IN"},
					Parent:   "DISTRIBUTOR1",
				},
			},
			want: &model.Distributor{
				Name:              "DISTRIBUTOR2",
				Includes:          []string{"IN"},
				Excludes:          []string{},
				ParentDistributor: "DISTRIBUTOR1",
			},
		},
	}
	utils.LoadCities("../cities.csv")
	utils.InitDistributors()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddDistributor(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddDistributor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDistributor() = %v, want %v", got, tt.want)
			}
		})
	}
}
