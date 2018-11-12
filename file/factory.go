package file

// Csv struct contains all the fields present as column in CSV file
type Csv struct {
	CityName     string
	ProvinceName string
	CountryName  string
}
type FileCsv struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}
type Distributor struct {
	Name       string
	ParentName string
	InList     []Csv
	Exlist     []Csv
}
