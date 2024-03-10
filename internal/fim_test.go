package internal

import (
	"reflect"
	"testing"
)

func TestAddFilm(t *testing.T) {
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
			if err := AddFilm(tt.args.id, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("AddFilm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidFilm(t *testing.T) {
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
			if got := IsValidFilm(tt.args.id); got != tt.want {
				t.Errorf("IsValidFilm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilmDB(t *testing.T) {
	tests := []struct {
		name string
		want map[FilmID]FilmName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilmDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilmDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveFilm(t *testing.T) {
	type args struct {
		id string
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
			if err := RemoveFilm(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFilm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
