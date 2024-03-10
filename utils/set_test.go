package utils

import (
	"reflect"
	"testing"
)

func TestNewSet(t *testing.T) {
	tests := []struct {
		name string
		want Set
	}{
		// TODO: Add test cases.
		{"Empty Set", Set{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		s    Set
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.item)
		})
	}
}

func TestSet_AddItems(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		s    Set
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.AddItems(tt.args.items)
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		s    Set
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	type args struct {
		set2 Set
	}
	tests := []struct {
		name string
		s    Set
		args args
		want Set
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersection(tt.args.set2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		s    Set
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.item)
		})
	}
}

func TestSet_Union(t *testing.T) {
	type args struct {
		set2 Set
	}
	tests := []struct {
		name string
		s    Set
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Union(tt.args.set2)
		})
	}
}
