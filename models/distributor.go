package models

type DistributorLocation struct {
	Country string
	State   string
	City    string
}

type Distributor struct {
	Name    string
	Include []DistributorLocation
	Exclude []DistributorLocation
	Parent  *Distributor
}
