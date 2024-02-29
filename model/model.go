package model

// Different structures used in the program

type Distributor struct {
	Name            string
	IncludeRegions  []string
	ExcludeRegions  []string
	SubDistributors []*Distributor // Children of a parent
}

type Location struct {
	Country string
	State   string
	City    string
}
