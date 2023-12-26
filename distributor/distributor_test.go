package distributor

import (
	"example.com/realimage_2016/constants"
	"example.com/realimage_2016/parsing"

	"testing"
)

// TestAddRegion tests the addRegion function
func TestAddRegion(t *testing.T) {
	// Updating the log file path
	constants.LogFilePath = "..\\realimage_2016.log"
	
	// Parsing the cities CSV file and storing the region data in a map
	parsing.ParseCsvFile("..\\cities.csv")

	// Checking if a valid region is added successfully
	include := make(map[string]map[string][]string)
	addRegion("TN-IN", include)
	if _, ok := include["IN"] ; ok {
		if _, ok := include["IN"]["TN"]; !ok {
			t.Errorf("Expected map with country code and province code, but province code not found")
		}
	} else {
		t.Errorf("Expected map with country code and province code, but country code not found")	
	}

	// Checking if a invalid region is ignored
	exclude := make(map[string]map[string][]string)
	addRegion("TTN-IN", include)
	if len(exclude) != 0 {
		t.Errorf("Expected a empty map since a invalid region is provided, but the region is added to map")
	}
}

// TestAddDistributor function tests the AddDistributor function
func TestAddDistributor(t *testing.T) {

	// Parsing the cities CSV file and storing the region data in a map
	parsing.ParseCsvFile("..\\cities.csv")

	// Passing details of distributor and validating the struct
	test_dis1 := AddDistributor("test_dis1", []string{"TN-IN"}, []string{"ABIRM-TN-IN"}, nil)
	if test_dis1.Id != "test_dis1" || len(test_dis1.Include["IN"]) != 1 || len(test_dis1.Exclude["IN"])!= 1 || test_dis1.Parent != nil{
		t.Errorf("Expected the ditributor struct with provided data, but got invalid data")
	}
}

// TestIsAuthorized function tests the IsAuthorized function with different test cases
func TestIsAuthorized(t *testing.T) {

	// Parsing the cities CSV file and storing the region data in a map
	parsing.ParseCsvFile("..\\cities.csv")

	// Adding distributor details and validating their permission
	includeRegions := []string{"TN-US", "IN", "JZRAH-RK-AE"}
	excludeRegions := []string{"EBRIN-TN-US", "TN-IN"}
	test_dis1 := AddDistributor("test_dis1", includeRegions, excludeRegions, nil)

	includeRegions = []string{"COLUB-TN-US", "AN-IN", "TN-IN"}
	excludeRegions = []string{"ACHMP-AP-IN"}
	test_dis2 := AddDistributor("test_dis2", includeRegions, excludeRegions, &test_dis1)

	// Test case to validate the permission for un-authorized regions for distributor test_dis1 and test_dis2
	unAuthorizedRegions := []string{"EBRIN-TN-US", "AL-US", "FORTH-VA-US", "ALGPU-TN-IN"}
	for _, region := range unAuthorizedRegions {
		if IsAuthorized(test_dis1, region) {
			t.Errorf("Expected false as response for un authorized regions, but got true for ditributor %s", test_dis1.Id)
		}
		if IsAuthorized(test_dis2, region) {
			t.Errorf("Expected false as response for un authorized regions, but got true for ditributor %s", test_dis2.Id)
		} 
	}

	// Test case to validate the permission for authorized regions for distributor test_dis1
	authorizedRegions := []string{"COLIR-TN-US", "COLUB-TN-US", "NIADS-AN-IN"}
	for _, region := range authorizedRegions {
		if !IsAuthorized(test_dis1, region) {
			t.Errorf("Expected true as response for authorized regions, but got false for ditributor %s", test_dis1.Id)
		}
	}

	// Test case to validate the permission for authorized regions for distributor test_dis2
	authorizedRegions = []string{"COLUB-TN-US", "NIADS-AN-IN"}
	for _, region := range authorizedRegions {
		if !IsAuthorized(test_dis2, region) {
			t.Errorf("Expected true as response for authorized regions, but got false for ditributor %s", test_dis2.Id)
		}
	}
}