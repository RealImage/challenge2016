package constants

// Regions to be included and excluded by distributor
type Permissions struct {
	Include map[string]map[string][]string
	Exclude map[string]map[string][]string
}

//  Distributor with region permissions and parent distributor if any
type Distributor struct {
	Id string
	Permissions
	Parent *Distributor
}

// Variable to store all the regions from CSV file
// Example: {"Country Code": {"Province code": ["City Code 1", "City Code 2"]}}
var RegionData map[string]map[string][]string

// Path to the CSV file containing cities
var CsvFilePath = "cities.csv"

// Path to log file
var LogFilePath = "realimage_2016.log"