package viewmodels

/*Struct for sending data to new page*/
type NewDistributorVM struct {
	PageTitle		string
	UniqueCountries		[]string
	UniqueProvinces		[][]string
	AllCities		[][]string
	DistributorCities	map[string][][]string
}

/*Struct to send data to the list page*/
type ListDistributorVM struct {
	PageTitle		string
	List			[]string
}