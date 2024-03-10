package internal

import (
	"reflect"
	"testing"
)

func TestAddDistributor(t *testing.T) {
	type args struct {
		id   string
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddDistributor(tt.args.id, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("AddDistributor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidDistributor(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidDistributor(tt.args.id); got != tt.want {
				t.Errorf("IsValidDistributor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDistributorDB(t *testing.T) {
	tests := []struct {
		name string
		want map[DistributorID]DistributorName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDistributorDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDistributorDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDistributor(t *testing.T) {
	type args struct {
		distributorID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveDistributor(tt.args.distributorID); (err != nil) != tt.wantErr {
				t.Errorf("RemoveDistributor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
