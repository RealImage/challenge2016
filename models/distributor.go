package models

type Distributor struct{
	ID string

	Parent *Distributor
}

type Rule struct{
	Dist_Id string
	Included map[City]bool
	Excluded map[City]bool
}