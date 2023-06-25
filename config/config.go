package config

//Any state, city or country name available in the csv file can be given for check

var (
	MainDistributer = "Distributer1"                  // Main distributer name
	Include         = []string{"Karnataka", "Kerala"} // Location to be included
	Exclude         = []string{"Chennai"}             // Location to be excluded
	Check           = "Bengaluru"                     //Check for distributer athourization
	SubDistributer1 = "Distributer2"                  // Sub Distributer name
	SubInclude1     = []string{"Karnataka"}           // Location included from main distributer
	SubExclude1     = []string{"Mysuru"}              // Location removed from main dsitributer
	SubCheck1       = "Bengaluru"                     // Check for sub distributer's authorisation
)
