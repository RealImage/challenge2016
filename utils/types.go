package utils

type Region struct {
    Country string
    State   string
    City    string
}

type Permission struct {
    Included []Region
    Excluded []Region
}

type location struct {
	cityCode     string
	provinceCode string
	countryCode  string
	cityName     string
	provinceName string
	countryName  string
}

type NewDistributer struct {
	Name    string
	Include []string
	Exclude []string
	Check   string
	Sub     []NewDistributer
}
