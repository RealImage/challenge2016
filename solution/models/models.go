package models

// City Code,Province Code,Country Code,City Name,Province Name,Country Name
type City struct {
	CityCode     string `json:"city_code" avro:"city_code"`
	ProvinceCode string `json:"province_code" avro:"province_code"`
	CountryCode  string `json:"country_code" avro:"country_code"`
	CityName     string `json:"city_name" avro:"city_name"`
	ProvinceName string `json:"province_name" avro:"province_name"`
	CountryName  string `json:"country_name" avro:"country_name"`
}

// Permission represents the permissions granted to a distributor
type Permission struct {
	Includes string `json:"includes" avro:"includes"`
	Excludes string `json:"excludes" avro:"excludes"`
}

// Distributor represents a distributor
type Distributor struct {
	ID      int      `json:"id" avro:"id"`
	Name    string   `json:"name" avro:"name"`
	Include []string `json:"includes" avro:"includes"`
	Exclude []string `json:"excludes" avro:"excludes"`
	Parent  string   `json:"parent" avro:"parent"`
	Child   string   `json:"child" avro:"child"`
}
