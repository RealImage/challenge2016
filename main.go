package main

import (
	"fmt"

	"github.com/souvikhaldar/challenge2016/auxilary"
	"github.com/souvikhaldar/challenge2016/file"
)

func main() {
	var csvSlice []file.Csv
	go file.Readfile("cities.csv", &csvSlice)
	var t, in, ex int
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
		fmt.Println("Result of CheckInclusion ", dist.CheckInclusion(dists))
		fmt.Println("Result of checkExclusion ", dist.CheckExclusion(dists))
		if dist.CheckInclusion(dists) == "Fine" && dist.CheckExclusion(dists) == "Fine" {
			parent := file.GetParent(dist.ParentName, dists)
			dist.AppendExlist(parent)
		}
		dists = append(dists, dist)
	}
	fmt.Println("Entered details of the distributors are:")
	for _, d := range dists {
		fmt.Println(d)
	}
}
