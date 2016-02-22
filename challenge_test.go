package main

import (
	//"fmt"
	"testing"
)

func TestSingleLevelPermissionCheck(t *testing.T) {
	Load_rule_csv()
	if !HasAuthorized("D1", "IN") {
		t.Errorf("D1 has authorization to distribute in IN")
	}

	if !HasAuthorized("D1", "US") {
		t.Errorf("D1 has authorization to distribute in US")
	}

	if HasAuthorized("D1", "UK") {
		t.Errorf("D1 doesn't have authoriUKzation to distribute in UK")
	}

	if !HasAuthorized("D2", "IN") {
		t.Errorf("D2 authorization to distribute in IN")
	}

	if !HasAuthorized("D2", "PUNE-MH-IN") {
		t.Errorf("D2 has authorization to distribute in MH-IN")
	}

	if HasAuthorized("D3", "HUBLI-KA-IN") {
		t.Errorf("D3 doesn't have authorization to distribute in HUBLI-KA-IN as D2 Doesn't hve it.")
	}
}

func TestMultipleLevelPermissionCheck(t *testing.T) {
	Load_rule_csv()

	if HasAuthorized("D2", "TN-IN") {
		t.Errorf("D2 doesn't have authorization to distribute in TN-IN")
	}

	if HasAuthorized("D2", "KA-IN") {
		t.Errorf("D2 doesn't have authorization to distribute in KA-IN")
	}

	if HasAuthorized("D2", "CENAI-TN-IN") {
		t.Errorf("D2 doesn't have authorization to distribute in KA-IN")
	}
}
