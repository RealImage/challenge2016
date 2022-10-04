package utils

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
	Sub     SubDistributer
}
type SubDistributer struct {
	Name    string
	Include []string
	Exclude []string
	Check string
}
