package main

import (
	"fmt"

	"github.com/souvikhaldar/challenge2016/auxilary"
	"github.com/souvikhaldar/challenge2016/file"
)

func main() {
	var csvSlice []file.Csv
	go file.Readfile("cities.csv", &csvSlice)
	var t, in, ex, n int
	var dists []file.Distributor
	fmt.Println("Enter the number of inputs: ")
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var dist file.Distributor
		fmt.Print("Enter name: ")
		fmt.Scanf("%s", &dist.Name)
		fmt.Print("Enter parent's name: ")
		fmt.Scanf("%s", &dist.ParentName)
		fmt.Print("Enter number of regions to include: ")
		fmt.Scanf("%d", &in)
		auxilary.FillSlice(in, &dist.InList)
		fmt.Print("Enter number of regions to exclude: ")
		fmt.Scanf("%d", &ex)
		auxilary.FillSlice(ex, &dist.Exlist)

		if file.CheckInclusion(dist.ParentName, dist, dists) == "Fine" && file.CheckExclusion(dist.ParentName, dist, dists) == "Fine" {
			parent := file.GetParent(dist.ParentName, dists)
			dist.AppendExlist(parent)
		} else {
			file.CheckInclusion(dist.ParentName, dist, dists)
			file.CheckExclusion(dist.ParentName, dist, dists)
		}
		dists = append(dists, dist)
	}
	fmt.Println("Enter number of regions to check: ")
	fmt.Scanf("%d", &n)
	for l := 0; l < n; l++ {
		var regions file.Distributor
		var name string
		fmt.Print("Enter name: ")
		fmt.Scanf("%s", &name)
		auxilary.FillSlice(1, &regions.InList)
		c := file.CheckInclusion(name, regions, dists)
		d := file.CheckExclusion(name, regions, dists)
		if c == "Fine" && d == "Fine" {
			fmt.Printf("YES! %s has permissions \n", name)
		} else {
			fmt.Printf("NO! %s doesn't have permission: ", c, d)
		}

	}

}
