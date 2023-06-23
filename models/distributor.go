package models

type Distributor struct{
	Name string `json:"name"`
	Include []Location `json:"include"`
	Exclude []Location  `json:"exclude"`
	ParentDistributor *string `json:"parentDistributor"`
}


type Location struct{
	City string `json:"city"`
	CityCode string `json:"cityCode"`
	Province string `json:"province"`
	ProvinceCode string `json:"provinceCode"`
	Country string `json:"country"`
	CountryCode string `json:"countryCode"`
}

type CheckPermission struct{
	DistributorName *string `json:"distributorName"`
	Loc *Location           `json:"loc"`
}