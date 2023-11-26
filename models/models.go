package models

// City  type store City Info
type City struct {
	CityCode     string `json:"cityCode"`
	ProvinceCode string `json:"provinceCode"`
	CountryCode  string `json:"countryCode"`
	CityName     string `json:"cityName"`
	ProvinceName string `json:"provinceName"`
	CountryName  string `json:"countryName"`
}

type PermissionType string

const (
	CountryType PermissionType = "country"
	StateType   PermissionType = "state"
	CityType    PermissionType = "city"
)

type AuthType string

const (
	Include AuthType = "include"
	Exclude AuthType = "exclude"
)

type Permission struct {
	IncludeMap map[string]PermissionType
	ExcludeMap map[string]PermissionType
}

//type PermissionMap map[string]Permission
type DistributerMap map[string]Permission //Dname- permission

type Distributer struct {
	Name            string
	InputPermission string
	AuthType        AuthType
}

// CountryMap store the key value pair of country name and stateMap of the states in that country
type CountryMap map[string]StateMap

// CityMap store key value pair of city name and City Info  of the same city
type CityMap map[string]City

// StateMap store key value pair of State Name and CityMap of cities coming in that state
type StateMap map[string]CityMap

type Reader interface {
	MakeDataStore()
}
type Operation interface {
	AddDistributer(countryMap CountryMap)
}
type InputModel struct{
	Name string
	Permission string
	AuthType AuthType
}
