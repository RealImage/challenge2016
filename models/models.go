package models

type Permission struct {
	Country string
	State   string
	City    string
}

type Distributor struct {
	Name    string
	Include Permission
	Exclude Permission
	Parent  *Distributor
}
