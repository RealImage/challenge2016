package util

type Authority struct {
	Name         string
	Distributors map[string]*Distributor
}

type Distributor struct {
	Name       string
	Permission Permissions
}
type Permissions struct {
	CountryPermission  CountryPermission
	ProvincePermission ProvincePermission
}

type CountryPermission struct {
	Allowed map[Country]ProvincePermission
}

type ProvincePermission struct {
	NotAllowed map[Province]int
	Allowed    map[Province]CityPermission
}

type CityPermission struct {
	Allowed    map[City]int
	NotAllowed map[City]int
}
type Country struct {
	Name string
	Code string
}

type Province struct {
	Name string
	Code string
}

type City struct {
	Name string
	Code string
}
