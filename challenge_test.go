package main

import (
	"testing"
)

func TestSingleDistributerPermission(t *testing.T) {
	Load_rule_csv()
	//PrintDistributerMap(DistributerMap)
	if !CheckPermission("D1", "IN") {
		t.Errorf("D1 has authorization to distribute in IN")
	}

	if !CheckPermission("D1", "US") {
		t.Errorf("D1 has authorization to distribute in US")
	}

	if CheckPermission("D1", "UK") {
		t.Errorf("D1 doesn't have authoriUKzation to distribute in UK")
	}

	if !CheckPermission("D2", "IN") {
		t.Errorf("D2 authorization to distribute in IN")
	}

	//if CheckPermission("D2", "MH-IN") {
	//	t.Errorf("D2 authorization to distribute in MH-IN")
	//}

	if !CheckPermission("D3", "HUBLI-KA-IN") {
		t.Errorf("D3 has authorization to distribute in HUBLI-KA-IN")
	}
}
