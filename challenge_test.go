package main

import (
	"testing"
)

func TestSingleLevelPermissionCheck(t *testing.T) {
	LoadRuleCsv()
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

	// As when checking only we are identifying invalid authorization mapping.
	// Could be added when creating the data structure itself and remove the
	// conflicting distributer and throw error.
	if HasAuthorized("D3", "HUBLI-KA-IN") {
		t.Errorf("D3 doesn't have authorization to distribute in HUBLI-KA-IN as D2 Doesn't hve it.")
	}
}

func TestMultipleLevelPermissionCheck(t *testing.T) {
	LoadRuleCsv()

	if HasAuthorized("D2", "TN-IN") {
		t.Errorf("D2 doesn't have authorization to distribute in TN-IN")
	}

	if HasAuthorized("D2", "KA-IN") {
		t.Errorf("D2 doesn't have authorization to distribute in KA-IN")
	}

	if HasAuthorized("D2", "CENAI-TN-IN") {
		t.Errorf("D2 doesn't have authorization to distribute in KA-IN")
	}

	if HasAuthorized("D3", "TN-IN") {
		t.Errorf("D3 doesn't have authorization to distribute in TN-IN")
	}

	if HasAuthorized("D3", "KA-IN") {
		t.Errorf("D3 doesn't have authorization to distribute in KA-IN")
	}

	if !HasAuthorized("D3", "MH-IN") {
		t.Errorf("D3 has authorization to distribute in MH-IN")
	}

	if !HasAuthorized("D3", "PUNE-MH-IN") {
		t.Errorf("D3 has authorization to distribute in PUNE-MH-IN")
	}
	if HasAuthorized("D3", "CN") {
		t.Errorf("D3 has doesn't have authorization to distribute in CN")
	}

	if HasAuthorized("D4", "KA-IN") {
		t.Errorf("D4 has doesn't have authorization to distribute in KA-IN")
	}

	if HasAuthorized("D4", "KL-IN") {
		t.Errorf("D4 doesn't have authorization to distribute in KL-IN")
	}

	if !HasAuthorized("D4", "PUNE-MH-IN") {
		t.Errorf("D4 has authorization to distribute in PUNE-MH-IN")
	}

	if !HasAuthorized("D5", "IN") {
		t.Errorf("D5 authorization to distribute in IN")
	}

	if HasAuthorized("D5", "PY-IN") {
		t.Errorf("D5 doesn't have authorization to distribute in PY-IN")
	}

	if HasAuthorized("D5", "KA-IN") {
		t.Errorf("D5 doesn't have authorization to distribute in KA-IN")
	}
	if HasAuthorized("D5", "KL-IN") {
		t.Errorf("D5 doesn't have authorization to distribute in KL-IN")
	}

	if HasAuthorized("D6", "KL-IN") {
		t.Errorf("D6 doesn't have authorization to distribute in KL-IN")
	}
	if HasAuthorized("D6", "PY-IN") {
		t.Errorf("D6 doesn't have authorization to distribute in PY-IN")
	}
	if !HasAuthorized("D6", "TN-IN") {
		t.Errorf("D6 has authorization to distribute in TN-IN")
	}

	if HasAuthorized("D7", "TN-IN") {
		t.Errorf("D7 has authorization to distribute in TN-IN")
	}

	if HasAuthorized("D7", "HUBLI-KA-IN") {
		t.Errorf("D7 doesn't have authorization to distribute in HUBLI-KA-IN as D1 doesn't have authorization in KA-IN")
	}
	if HasAuthorized("D8", "HUBLI-KA-IN") {
		t.Errorf("D7 has authorization to distribute in TN-IN")
	}

	if HasAuthorized("D10", "KL-IN") {
		t.Errorf("D10 doesn't have authroization to distribute in KL-IN")
	}
	if HasAuthorized("D10", "NY-US") {
		t.Errorf("D10 doesn't have authroization to distribute in NY-US")
	}
	if HasAuthorized("D10", "KA-IN") {
		t.Errorf("D10 doesn't have authroization to distribute in KA-IN")
	}
	if HasAuthorized("D10", "CENAI-TN-IN") {
		t.Errorf("D10 doesn't have authroization to distribute in CENAI-TN-IN")
	}
	if !HasAuthorized("D10", "MH-IN") {
		t.Errorf("D10 have authroization to distribute in MH-IN")
	}
	if !HasAuthorized("D10", "DC-US") {
		t.Errorf("D10 have authroization to distribute in DC-US")
	}

	// Only Exclusion distributers
	if HasAuthorized("D11", "KL-IN") {
		t.Errorf("D11 doesn't have authroization to distribute in KL-IN")
	}
	if HasAuthorized("D11", "TN-IN") {
		t.Errorf("D11 doesn't have authroization to distribute in TN-IN")
	}
	if HasAuthorized("D11", "AH-DC-US") {
		t.Errorf("D11 doesn't have authroization to distribute in AH-DC-US")
	}
	if !HasAuthorized("D11", "MH-IN") {
		t.Errorf("D11 have authroization to distribute in MH-IN")
	}
	if !HasAuthorized("D11", "DC-IN") {
		t.Errorf("D11 have authroization to distribute in DC-IN")
	}
	if !HasAuthorized("D11", "IN") {
		t.Errorf("D11 have authroization to distribute in IN")
	}
	if !HasAuthorized("D11", "US") {
		t.Errorf("D11 have authroization to distribute all the places in US")
	}
}
