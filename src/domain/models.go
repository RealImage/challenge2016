package domain

var CountryMap map[string][][]string
var DistributorMap map[string]Distributor

type Distributor struct {
	Name               string
	IncludePermissions []string
	ExcludePermissions []string
	ParentDistributor  *Distributor
}
