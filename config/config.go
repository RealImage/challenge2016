package config

//Any state, city or country name available in the csv file can be given for check

var (
	MainDistributer = "Distributer1"        // Main distributer name
	Include         = []string{"India", "China"}     // Location to be included
	Exclude         = []string{"Karnataka"} // Location to be excluded
	Check           = "Chennai"             //Check for distributer athourization
	SubDistributer  = "Distributer2"        // Sub Distributer name
	SubInclude      = []string{"Kerala"} // Location included from main distributer
	SubExclude      = []string{"Cochin"} // Location removed from main dsitributer
	SubCheck        = "Cochin"           // Check for sub distributer's authorisation
)
