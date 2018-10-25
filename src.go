package main

import (
	"container/list"
	//"encoding/csv"
	"fmt"
	//"log"
	//"strings"
	//"os"
)

type ByPass struct {
	complete bool
	ptr      *list.List
}

/*
type Distributor struct {
	countryList  *list.List
	provinceList *list.List
	cityList     *list.List
	byPass       [3]*ByPass
}
*/
type Distributor struct {
	countryList  *list.List
	provinceList *list.List
	cityList     *list.List
	byPass       []ByPass
}

func main() {
	var nd int
	fmt.Println("enter no. of distributors")
	fmt.Scanf("%d", &nd)
	distrib := make([]Distributor, nd)
var d,m int
	for d = 0; d < nd; d++ {
		distrib[d].countryList = list.New()
		distrib[d].provinceList = list.New()
		distrib[d].cityList = list.New()

		distrib[d].byPass = make([]ByPass, 3)

		distrib[d].byPass[0].complete = false
		distrib[d].byPass[1].complete = false
		distrib[d].byPass[2].complete = false
		distrib[d].byPass[0].ptr = distrib[d].countryList
		distrib[d].byPass[1].ptr = distrib[d].provinceList
		distrib[d].byPass[2].ptr = distrib[d].cityList

		var numCountries, numProvinces, numCities int
		var countries, provinces, cities string

		fmt.Println("enter no. of countries")
		fmt.Scanf("%d", &numCountries)
fmt.Println("no.of countries:",numCountries)
		for i := 0; i < numCountries; i++ {
			fmt.Scanf("%s", &countries)
			distrib[d].countryList.PushBack(countries)
		}
		fmt.Println("enter no. of provinces")
		fmt.Scanf("%d", &numProvinces)
fmt.Println("no.of provinces:",numProvinces)

		for i := 0; i < numProvinces; i++ {
			fmt.Scanf("%s", &provinces)
			distrib[d].provinceList.PushBack(provinces)
		}
		fmt.Println("enter no. of cities")
		fmt.Scanf("%d", &numCities)
fmt.Println("no.of cities:",numCities)

		for i := 0; i < numCities; i++ {
			fmt.Scanf("%s", &cities)
			distrib[d].cityList.PushBack(cities)
		}

		fmt.Println("\n\n")
		for e := distrib[d].byPass[0].ptr.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value.(string))
		}
		fmt.Println("\n\n")
		for e := distrib[d].byPass[1].ptr.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value.(string))
		}
		fmt.Println("\n\n")
		for e := distrib[d].byPass[2].ptr.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value.(string))
		}
	}
	d0, m0 := getOutput(distrib, d, m, 0)
	if d0 == m0 {
		d1, m1 := getOutput(distrib, d, m, 1)
		if d1 == m1 {
			d2, m2 := getOutput(distrib, d, m, 2)
			if d2 == m2 {
				fmt.Println("\nRules followed")
			} else {
				fmt.Println("\nRules not followed")
			}
		}
	}

}
func getOutput(distrib []Distributor, d int, m int, level int) (string, string) {
var ea, eb *list.Element
	if level == 0 {
		ea = distrib[d].countryList.Front()
		eb = distrib[m].countryList.Front()
	} else if level == 1 {
		ea = distrib[d].provinceList.Front()
		eb = distrib[m].provinceList.Front()
	} else {
		ea = distrib[d].cityList.Front()
		eb = distrib[m].cityList.Front()
	}

	for ; ea != nil; ea = ea.Next() {
		for ; eb != nil; eb = eb.Next() {

			if ea.Value.(string) == eb.Value.(string) {
				fmt.Println("\nyes same ", ea.Value.(string), " ", eb.Value.(string))
				return ea.Value.(string), eb.Value.(string)
			} else {
				fmt.Println("\nnot same", ea.Value.(string), " ", eb.Value.(string))
				return ea.Value.(string), eb.Value.(string)
			}
		}
	}
return "",""
}
