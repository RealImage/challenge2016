package viewmodels

type NewDistributorVM struct {
	PageTitle		string
	UniqueCountries		[]string
	UniqueProvinces		[][]string
	AllCities		[][]string
	DistributorCities	map[string][][]string
}


type ListDistributorVM struct {
	PageTitle		string
	List			[]string
}