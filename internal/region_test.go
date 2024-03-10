package internal

import (
	"reflect"
	"testing"
)

func TestAddRegion(t *testing.T) {
	type args struct {
		cityID      string
		stateID     string
		countryID   string
		cityName    string
		stateName   string
		countryName string
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
			if err := AddRegion(tt.args.cityID, tt.args.stateID, tt.args.countryID, tt.args.cityName, tt.args.stateName, tt.args.countryName); (err != nil) != tt.wantErr {
				t.Errorf("AddRegion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidRegion(t *testing.T) {
	type args struct {
		regionID string
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
			if got := IsValidRegion(tt.args.regionID); got != tt.want {
				t.Errorf("IsValidRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRegionDB(t *testing.T) {
	tests := []struct {
		name string
		want map[CountryID]map[StateID]map[CityID]RegionData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegionDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegionDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadRegionDataFromLocalCSV(t *testing.T) {
	type args struct {
		path string
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
			if err := ReadRegionDataFromLocalCSV(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("ReadRegionDataFromLocalCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadRegionDataFromRemoteCSV(t *testing.T) {
	type args struct {
		url string
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
			if err := ReadRegionDataFromRemoteCSV(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("ReadRegionDataFromRemoteCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isValidCity(t *testing.T) {
	type args struct {
		id     string
		cityDB map[CityID]RegionData
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
			if got := isValidCity(tt.args.id, tt.args.cityDB); got != tt.want {
				t.Errorf("isValidCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidCountry(t *testing.T) {
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
			if got := isValidCountry(tt.args.id); got != tt.want {
				t.Errorf("isValidCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidState(t *testing.T) {
	type args struct {
		id      string
		stateDB map[StateID]map[CityID]RegionData
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
			if got := isValidState(tt.args.id, tt.args.stateDB); got != tt.want {
				t.Errorf("isValidState() = %v, want %v", got, tt.want)
			}
		})
	}
}
