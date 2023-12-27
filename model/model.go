package model

// Permissions for distributors
type Permissions struct {
	Include          []string                `json:"include"`
	Exclude          []string                `json:"exclude"`
	ChildPermissions map[string]*Permissions `json:"childPermissions,omitempty"`
}

// Distribution details
type Distribution struct {
	Data               map[string]Permissions `json:"data"`
	NestedDistribution *Distribution
}
type DistributionReq struct {
	Name        string      `json:"distributor"`
	Permissions Permissions `json:"permissions"`
}
