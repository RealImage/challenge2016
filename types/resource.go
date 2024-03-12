package types

type (
	Distributor struct {
		Name     string
		Parent   *Distributor
		Includes []LocationIdentifier
		Excludes []LocationIdentifier
	}

	LocationIdentifier struct {
		CountryCode  string
		ProvinceCode string
		CityCode     string
		CountryName  string
		ProvinceName string
		CityName     string
	}

	DistributorRequest struct {
		DistributorName       string     `json:"distributorName,omitempty"`
		ParentDistributorName string     `json:"parentDistributorName"`
		Includes              []Location `json:"includes"`
		Excludes              []Location `json:"excludes"`
	}

	Location struct {
		CountryCode  string `json:"countryCode"`
		ProvinceCode string `json:"provinceCode"`
		CityCode     string `json:"cityCode"`
	}

	LocationDetailsReq struct {
		DistributorName string
		CountryCode     string
		ProvinceCode    string
		CityCode        string
	}

	GenericResponse struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
