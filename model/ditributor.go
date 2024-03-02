package model

type Distributor struct {
	Name              string
	ParentDistributor string
	Includes          []string
	Excludes          []string
}
