package dto

type City struct {
	Name string `json:"name"`
}

type State struct {
	Name   string `json:"state"`
	Cities []City `json:"cities"`
}

type Country struct {
	Name   string  `json:"country"`
	States []State `json:"states"`
}

type Location struct {
	CityCode     string `csv:"City Code"`
	ProvinceCode string `csv:"Province Code"`
	CountryCode  string `csv:"Country Code"`
	CityName     string `csv:"City Name"`
	ProvinceName string `csv:"Province Name"`
	CountryName  string `csv:"Country Name"`
}
