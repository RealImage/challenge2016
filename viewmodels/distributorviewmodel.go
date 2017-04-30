package viewmodels

/*Struct for sending data to new page*/
/*type NewDistributorVM struct {
	PageTitle		string
	UniqueCountries		[]string
	UniqueProvinces		[][]string
	AllCities		[][]string
	DistributorCities	map[string][][]string
}*/

type NewDistributorVM struct {
	PageTitle		string
	UniqueCountries		[][]string
	UniqueProvinces		[][]string
	AllCities		[][]string
	DistributorCities	map[string][][]string
	DistributorCountries	map[string][][]string
	DistributorNames	[]string
}

/*Struct to send data to the list page*/
type ListDistributorVM struct {
	PageTitle		string
	List			[]string
}

/*Struct to send data to the view page*/
type ViewDistributorVM struct {
	PageTitle		string
	AllCities		[][]string
}