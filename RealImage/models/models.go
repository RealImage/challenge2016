package models

// Permissions for distributors
type Permission struct {
	Include []Location `json:"Include,omitempty"`
	Exclude []Location `json:"Exclude,omitempty"`
}

// Location represents a location with its details.
type Location struct {
	CityCode     string `json:"cityCode,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
	CountryCode  string `json:"countryCode"`
	CityName     string `json:"cityName,omitempty"`
	ProvinceName string `json:"provinceName,omitempty"`
	CountryName  string `json:"countryName"`
}

// Input inquiry to be filled by the user in the JSON
type EnquiryForm struct {
	DistributorName string   `json:"DistributorName"`
	Location        Location `json:"Location"`
}
