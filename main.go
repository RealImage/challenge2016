package main

import (
	"fmt"

	"github.com/souvikhaldar/challenge2016/auxilary"
	"github.com/souvikhaldar/challenge2016/file"
)

func main() {
	var csvSlice []file.FileCsv
	go file.Readfile("cities.csv", &csvSlice)
	var t, in, ex, n int
	var dists []file.Distributor
	fmt.Println("Enter the number of distributors: ")
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Printf("%d) \n", i+1)
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
		a := file.CheckInclusion(dist.ParentName, dist, dists)
		b := file.CheckExclusion(dist.ParentName, dist, dists)
		if a == "Fine" && b == "Fine" {
			if dist.ParentName != "" {
				fmt.Printf("%s can distrubute work to %s \n", dist.ParentName, dist.Name)
			}
			parent := file.GetParent(dist.ParentName, dists)
			dist.AppendExlist(parent)
			dists = append(dists, dist)
		} else {
			if a == "Fine" && b != "Fine" {
				fmt.Println("Permission Error! Child can't include regions that is excluded in parent: ", b)
			} else {
				fmt.Println("Permission Error! Child can't include regions that is not included in the parent: ", a)
			}
		}

	}
	fmt.Println("Enter number of regions to check: ")
	fmt.Scanf("%d", &n)
	for l := 0; l < n; l++ {
		var regions file.Distributor
		var name string
		fmt.Println("Enter name: ")
		fmt.Scanf("%s", &name)
		auxilary.FillSlice(1, &regions.InList)
		c := file.CheckInclusion(name, regions, dists)
		d := file.CheckExclusion(name, regions, dists)
		if c == "Fine" && d == "Fine" {
			fmt.Printf("YES. %s has permissions.\n", name)
		} else if c == "Fine" && d != "Fine" {
			fmt.Printf("NO. %s doesn't have required permissions. %s \n", name, d)
		} else if c != "Fine" && d == "Fine" {
			fmt.Printf("NO. %s doesn't have required permissions. %s \n", name, c)
		} else {
			fmt.Printf("NO. %s doesn't have required permissions.\n", name)
		}

	}

}
