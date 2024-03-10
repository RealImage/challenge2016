package internal

import (
	"reflect"
	"testing"
)

func TestAuthorizeDistributor(t *testing.T) {
	type args struct {
		filmID             string
		ownerDistributorID string
		agentDistributorID string
		includes           []string
		excludes           []string
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
			if err := AuthorizeDistributor(tt.args.filmID, tt.args.ownerDistributorID, tt.args.agentDistributorID, tt.args.includes, tt.args.excludes); (err != nil) != tt.wantErr {
				t.Errorf("AuthorizeDistributor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllFilmsAuthorizedForDistributor(t *testing.T) {
	type args struct {
		distributorID string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllFilmsAuthorizedForDistributor(tt.args.distributorID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllFilmsAuthorizedForDistributor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPermission(t *testing.T) {
	type args struct {
		filmID        string
		distributorID string
		region        string
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
			if got := HasPermission(tt.args.filmID, tt.args.distributorID, tt.args.region); got != tt.want {
				t.Errorf("HasPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthorizationDB(t *testing.T) {
	tests := []struct {
		name string
		want map[FilmID]map[DistributorID]AuthorizationMetaData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthorizationDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthorizationDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveAuthorizationForDistributor(t *testing.T) {
	type args struct {
		filmID        string
		distributorID string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveAuthorizationForDistributor(tt.args.filmID, tt.args.distributorID)
		})
	}
}

func TestRemoveAuthorizationForFilm(t *testing.T) {
	type args struct {
		filmID string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveAuthorizationForFilm(tt.args.filmID)
		})
	}
}
