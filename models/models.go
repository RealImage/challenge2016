package models

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

type DistributerMap map[string]Permission

type Distributer struct {
	Name            string
	InputPermission string
	AuthType        AuthType
}

type CountryMap map[string]StateMap

type CityMap map[string]City

type StateMap map[string]CityMap

type Reader interface {
	MakeDataStore()
}
type Operation interface {
	AddDistributer(countryMap CountryMap)
}
type InputModel struct {
	Name       string
	Permission string
	AuthType   AuthType
}
