package models

type Distributor struct{
	Included map[City]bool
	Excluded map[City]bool

	Parent *Distributor
	Sub []*Distributor
}