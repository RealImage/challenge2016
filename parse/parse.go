package parse

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ParseCSV parses the CSV file and returns a map representing the region structure.
func ParseCSV(filePath string) (map[string]map[string]map[string]bool, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Parse the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}

	// Create a map to store regions
	regions := make(map[string]map[string]map[string]bool)

	// Iterate through the CSV records and populate the regions map
	for _, record := range records {
		cityCode := record[0]
		provinceCode := record[1]
		regionCode := record[2]

		// Check if the region exists in the map
		region, regionExists := regions[regionCode]
		if !regionExists {
			// If the region doesn't exist, create a new entry
			region = make(map[string]map[string]bool)
			regions[regionCode] = region
		}

		// Check if the province exists in the region map
		province, provinceExists := region[provinceCode]
		if !provinceExists {
			// If the province doesn't exist, create a new entry
			province = make(map[string]bool)
			region[provinceCode] = province
		}

		// Add the city to the province
		province[cityCode] = true
	}

	return regions, nil
}