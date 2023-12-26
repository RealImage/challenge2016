package parsing

import (
	"example.com/realimage_2016/constants"

	"testing"
)

// TestParseCsvFile function tests the functionality of ParseCsvFile
func TestParseCsvFile(t *testing.T) {

	// Updating the log file path
	constants.LogFilePath = "..\\realimage_2016.log"
	
	// Giving an invalid path
	if ParseCsvFile("sample.csv") {
		t.Errorf("Given a invalid csv file path, expected false as a response, but got true")
	}

	// Checking a random data by providing the right path
	ParseCsvFile("..\\cities.csv")
	if provinces, ok := constants.RegionData["IN"]; ok {
		if cities, ok := provinces["TN"]; ok {
			if len(cities) == 0 {
				t.Errorf("Expected parsed data from valid CSV file, but data not found")
			}
		} else {
			t.Errorf("Expected parsed data from valid CSV file, but data not found")
		}
	} else {
		t.Errorf("Expected parsed data from valid CSV file, but data not found")
	}
}