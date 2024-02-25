package common

type Response struct {
	Msg    string      `json:"_msg"`
	Status int         `json:"_status"`
	Data   interface{} `json:"data"`
}

type LocationIdentifier struct {
	CountryCode  string
	ProvinceCode string
	CityCode     string
	CountryName  string
	ProvinceName string
	CityName     string
}

func NewLocationIdentifier(countryCode, provinceCode, cityCode, countryName, provinceName, cityName string) LocationIdentifier {
	return LocationIdentifier{
		CountryCode:  countryCode,
		ProvinceCode: provinceCode,
		CityCode:     cityCode,
		CountryName:  countryName,
		ProvinceName: provinceName,
		CityName:     cityName,
	}
}

type Distributor struct {
	Name     string
	Parent   *Distributor
	Includes []LocationIdentifier
	Excludes []LocationIdentifier
}

var DistributorsMap = make(map[string]*Distributor)

func NewDistributor(name, parentName string) *Distributor {
	var parent *Distributor
	if parentName != "" {
		parent = DistributorsMap[parentName]

	}
	dist := Distributor{
		Name:     name,
		Parent:   parent,
		Includes: make([]LocationIdentifier, 0),
		Excludes: make([]LocationIdentifier, 0),
	}
	DistributorsMap[name] = &dist
	return &dist
}

type Location struct {
	CountryCode  string
	ProvinceCode string
	CityCode     string
}

type DistributorInput struct {
	Distributorname       string
	ParentDistributorName string
	Includes              []Location
	Excludes              []Location
}

var CountryMap = make(map[string]LocationIdentifier)
var ProvinceMap = make(map[string]LocationIdentifier)
var CityMap = make(map[string]LocationIdentifier)

type LocationAccessInput struct {
	Distributorname string
	CountryCode     string
	ProvinceCode    string
	CityCode        string
}
