package dtos

type Location struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type AuthorisationReq struct {
	DistributorName string      `json:"name"`
	Include         []*Location `json:"included"`
	Exclude         []*Location `json:"excluded"`
	ParentName      string      `json:"parent_name"`
}

type Distributor struct {
	Parent   string
	Included map[string]*Country
	Excluded map[string]*Country
}

type GetStatusReq struct {
	DistributorName string   `json:"name"`
	Region          Location `json:"region"`
}
