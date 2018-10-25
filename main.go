package main

import (
	"fmt"

	"github.com/souvikhaldar/challenge2016/auxilary"
	"github.com/souvikhaldar/challenge2016/file"
)

func main() {
	go file.Readfile("cities.csv")
	fmt.Println("Enter the number of inputs: ")
	var t, in, ex int
	var name, parentName string
	var inList, exList []file.Csv
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Print("Enter name: ")
		fmt.Scanf("%s", &name)
		fmt.Print("Enter parent's name: ")
		fmt.Scanf("%s", &parentName)
		fmt.Print("Enter number of regions to include: ")
		fmt.Scanf("%d", &in)
		auxilary.FillSlice(in, &inList)
		fmt.Print("Enter number of regions to exclude: ")
		fmt.Scanf("%d", &ex)
		auxilary.FillSlice(ex, &exList)
		fmt.Println("InList: ", inList)
		fmt.Println("ExList: ", exList)
	}
}
