package utils

import "github.com/saurabh-sde/challenge2016_saurabh/model"

type NewDistributorRequest struct {
	Name     string   `json:"name"`
	Includes []string `json:"includes,omitempty"`
	Excludes []string `json:"excludes,omitempty"`
	Parent   string   `json:"parent,omitempty"`
}

type CheckDistributorPermissionRequest struct {
	Name      string   `json:"name"`
	Locations []string `json:"locations"`
}

var DistributorMap map[string]model.Distributor

func InitDistributors() {
	DistributorMap = make(map[string]model.Distributor)
}
