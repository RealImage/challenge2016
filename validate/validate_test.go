package validate

import (
	"example.com/realimage_2016/constants"
	"example.com/realimage_2016/parsing"
	
	"testing"
)

// TestIsValidRegion function tests the functionality of IsValidRegion
func TestIsValidRegion(t *testing.T) {

	// Updating the log file path
	constants.LogFilePath = "..\\realimage_2016.log"

	// Parsing the cities CSV file and storing the region data in a map
	parsing.ParseCsvFile("..\\cities.csv")


	// Giving a valid region and checking the response
	if !IsValidRegion("TN-IN") {
		t.Errorf("Expected true as a response, but got false for valid region")
	}

	// Giving an invalid region and checking the response
	if IsValidRegion("NOO-TN-IN") {
		t.Errorf("Expected false as a response, but got true for invalid region")
	}
}