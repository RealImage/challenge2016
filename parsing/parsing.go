// Parse the CSV file and store the required details
package parsing

import (
	"example.com/realimage_2016/constants"
	"example.com/realimage_2016/logger"

	"encoding/csv"
	"os"
)

// ParseCsvFile function reads region data from a CSV file and stores it inside a map
// filePath - Path of CSV file which has the data of all cities in world
func ParseCsvFile(filePath string) bool{

	// To store logs inside log file
	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer log.Close()

	regionData := make(map[string]map[string][]string)

	// Opening the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Log("Error while opening file.", err)
		return false
	}
	defer file.Close()
	// Reading the contents of the file
	reader := csv.NewReader(file)
	fileContent, err := reader.ReadAll()
	if err != nil {
		log.Log("Error while reading CSV:", err)
		return false
	}
	// Parsing the contents of the CSV file and storing the required data
	for _, line := range fileContent[1:] {
		// Checking if the line has atleast 3 entries
		if len(line) > 3 {
			if _, ok := regionData[line[2]]; ok {
				if _, ok := regionData[line[2]][line[1]]; ok {
					regionData[line[2]][line[1]] = append(regionData[line[2]][line[1]], line[0])
				} else {
					regionData[line[2]][line[1]] = []string{line[0]}
				}
			} else {
				regionData[line[2]] = make(map[string][]string)
				regionData[line[2]][line[1]] = []string{line[0]}
			}
		}
	}

	constants.RegionData = regionData
	return true
}