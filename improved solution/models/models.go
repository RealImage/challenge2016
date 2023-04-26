package models

// Distributor represents a distributor
type Distributor struct {
	ID              int      `json:"id" avro:"id"`
	Name            string   `json:"name" avro:"name"`
	PermittedPlaces []string `json:"permitted_places" avro:"permitted_places"`
	Parent          string   `json:"parent" avro:"parent"`
	Child           string   `json:"child" avro:"child"`
}

// DistributorData represents the data for the distributor
type DistributorsModel struct {
	CountryStateMap    map[string][]string
	StateCityMap       map[string][]string
	CurrentDistributor Distributor
	Distributors       []Distributor
}