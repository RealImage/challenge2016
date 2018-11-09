package main

import (
	"fmt"
	"testing"

	"github.com/souvikhaldar/challenge2016/file"
)

var distributors, distributors2 []file.Distributor

func init() {
	distributors = []file.Distributor{
		{
			Name:       "dist1",
			ParentName: "",
			InList:     []file.Csv{{"", "", "india"}, {"", "", "us"}},
			Exlist:     []file.Csv{{"", "karnataka", "india"}, {"chennai", "tamilnadu", "india"}},
		},
		{
			Name:       "dist2",
			ParentName: "dist1",
			InList:     []file.Csv{{"", "", "india"}},
			Exlist:     []file.Csv{{"", "tamilnadu", "india"}},
		},
		{
			Name:       "dist3",
			ParentName: "dist2",
			InList:     []file.Csv{{"hubli", "karnataka", "india"}},
			Exlist:     []file.Csv{},
		},
	}
	distributors2 = []file.Distributor{
		{
			Name:       "dist1",
			ParentName: "",
			InList:     []file.Csv{{"", "", "india"}, {"", "", "us"}},
			Exlist:     []file.Csv{{"", "karnataka", "india"}, {"chennai", "tamilnadu", "india"}},
		},
		{
			Name:       "dist2",
			ParentName: "dist1",
			InList:     []file.Csv{{"", "", "india"}},
			Exlist:     []file.Csv{{"", "tamilnadu", "india"}},
		},
		{
			Name:       "dist3",
			ParentName: "dist2",
			InList:     []file.Csv{{"hubli", "karnataka", "india"}},
			Exlist:     []file.Csv{},
		},
	}
}
func TestDistAddition(t *testing.T) {
	var dist file.Distributor
	var n int
	for n, dist = range distributors {
		fmt.Println(n)
		parent := file.GetParent(dist.ParentName, distributors2)
		dist.AppendExlist(parent)
		if n == 1 {
			a := file.CheckInclusion(dist.ParentName, dist, distributors2)
			b := file.CheckExclusion(dist.ParentName, dist, distributors2)
			fmt.Println(a, b)
			if a != "Fine" || b != "Fine" {
				t.Error("Distributor should have ideally been added but didn't")
			}
		}
		if n == 2 {
			c := file.CheckInclusion(dist.ParentName, dist, distributors2)
			d := file.CheckExclusion(dist.ParentName, dist, distributors2)
			fmt.Println(c, d)
			if c == "Fine" || d == "Fine" {
				t.Error("Distributor should have ideally not been added but did.")
			}
		}

	}
}
