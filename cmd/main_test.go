package main

import (
	"bufio"
	"testing"
)

func Test_readFilmAuthorizationInput(t *testing.T) {
	type args struct {
		reader *bufio.Reader
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
			if err := readFilmAuthorizationInput(tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("readFilmAuthorizationInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readPermissionCheckInput(t *testing.T) {
	type args struct {
		reader *bufio.Reader
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
			if err := readPermissionCheckInput(tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("readPermissionCheckInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readRegionDataInput(t *testing.T) {
	type args struct {
		reader *bufio.Reader
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
			if err := readRegionDataInput(tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("readRegionDataInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
