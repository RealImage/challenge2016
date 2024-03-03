package model

// Distributor represents a distributor with its include and exclude regions
type Distributor struct {
	Name            string
	IncludeRegions  []string
	ExcludeRegions  []string
	SubDistributors []*Distributor // Children of a parent
}

// Location represents a geographical location
type Location struct {
	Country string
	State   string
	City    string
}

// Constants for slice indices
const (
	IncludeRegionsIndex = iota
	ExcludeRegionsIndex
)

// NewDistributor creates a new Distributor with specified include and exclude regions
func NewDistributor(name string, includeRegions, excludeRegions []string) *Distributor {
	return &Distributor{
		Name:            name,
		IncludeRegions:  includeRegions,
		ExcludeRegions:  excludeRegions,
		SubDistributors: make([]*Distributor, 0),
	}
}
