package main

type LocationData struct {
	AvailableLocations map[string]map[string]map[string]struct{}
	Distributors       map[string]DistributorData
	DistributorParent  map[string]string
}

type DistributorData struct {
	ParentDistributor string
	Included          map[string]struct{}
	Excluded          map[string]struct{}
}

// create distributor json body
type DistribRequestBody struct {
	Distributor       string   `json:distributor`
	ParentDistributor string   `json:parentdistributor`
	Include           []string `json:include`
	Exclude           []string `json:exclude`
}
