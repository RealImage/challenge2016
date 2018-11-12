package main

import (
	"testing"

	"github.com/souvikhaldar/challenge2016/file"
)

var distributors, distributors2 []file.Distributor

func init() {
	distributors = []file.Distributor{
		{
			Name:       "dist1",
			ParentName: "",
			InList:     []file.Csv{{CityName: "", ProvinceName: "", CountryName: "india"}, {CityName: "", ProvinceName: "", CountryName: "us"}},
			Exlist:     []file.Csv{{CityName: "", ProvinceName: "karnataka", CountryName: "india"}, {CityName: "chennai", ProvinceName: "tamilnadu", CountryName: "india"}},
		},
		{
			Name:       "dist2",
			ParentName: "dist1",
			InList:     []file.Csv{{CityName: "", ProvinceName: "", CountryName: "india"}},
			Exlist:     []file.Csv{{CityName: "", ProvinceName: "tamilnadu", CountryName: "india"}},
		},
		{
			Name:       "dist3",
			ParentName: "dist2",
			InList:     []file.Csv{{CityName: "hubli", ProvinceName: "karnataka", CountryName: "india"}},
			Exlist:     []file.Csv{},
		},
	}
	distributors2 = []file.Distributor{
		{
			Name:       "dist1",
			ParentName: "",
			InList:     []file.Csv{{CityName: "", ProvinceName: "", CountryName: "india"}, {CityName: "", ProvinceName: "", CountryName: "us"}},
			Exlist:     []file.Csv{{CityName: "", ProvinceName: "karnataka", CountryName: "india"}, {CityName: "chennai", ProvinceName: "tamilnadu", CountryName: "india"}},
		},
		{
			Name:       "dist2",
			ParentName: "dist1",
			InList:     []file.Csv{{CityName: "", ProvinceName: "", CountryName: "india"}},
			Exlist:     []file.Csv{{CityName: "", ProvinceName: "tamilnadu", CountryName: "india"}},
		},
	}
}

// TestDistAddition checks if one distributor can assign permission to another distributor
func TestDistAddition(t *testing.T) {
	var dist file.Distributor
	var n int
	for n, dist = range distributors {
		parent := file.GetParent(dist.ParentName, distributors2)
		dist.AppendExlist(parent)
		if n == 1 {
			a := file.CheckInclusion(dist.ParentName, dist, distributors2)
			b := file.CheckExclusion(dist.ParentName, dist, distributors2)
			if a != "Fine" || b != "Fine" {
				t.Error("Distributor should have ideally been added but didn't")
			}
		}

	}
}

// TestDistPermission checks whether a distributor has permission for the given region
func TestDistPermission(t *testing.T) {
	var regions file.Distributor
	regions.InList = []file.Csv{{
		CityName:     "chicago",
		ProvinceName: "illinois",
		CountryName:  "us",
	}}
	c := file.CheckInclusion("dist1", regions, distributors2)
	d := file.CheckExclusion("dist1", regions, distributors2)
	if c != "Fine" || d != "Fine" {
		t.Error("Distributor is supposed to have permission, but didn't.")
	}

}
