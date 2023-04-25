package handler

import "github.com/dunzoit/challenge2016/distributor"

// Distributor Define a struct for the distributor class
type distributors struct {
	Distributors []*distributor.Distributor
}

func (d distributors) AddDistributor(included, excluded []string) {
	newDistributor := distributor.NewDistributor(included, excluded)
	d.Distributors = append(d.Distributors, newDistributor)
	return
}
