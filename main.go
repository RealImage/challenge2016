package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"bufio"
	"strconv"
	"strings"
)

type Distributor struct {
	include [] Area
	exclude []Area
}

type Area struct {
	city string
	state string
	country string
}

type CityMap struct {
	CityName      string
	ProvinceName  string
	CountryName   string
}

type City struct {
	CityCode      string `csv:"City Code"`
	ProvinceCode  string `csv:"Province Code"`
	CountryCode   string `csv:"Country Code"`
	CityName      string `csv:"City Name"`
	ProvinceName  string `csv:"Province Name"`
	CountryName   string `csv:"Country Name"`
}

func main() {
	// Open the CSV file
	file, err := os.Open("cities.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Create a slice to hold the City structs
	var cities []City

	// Iterate over the records and populate the City structs
	for i, record := range records {
		if i == 0 {
			continue
		}
		city := City{
			CityCode:     strings.ToUpper(record[0]),
			ProvinceCode: strings.ToUpper(record[1]),
			CountryCode:  strings.ToUpper(record[2]),
			CityName:     strings.ToUpper(record[3]),
			ProvinceName: strings.ToUpper(record[4]),
			CountryName:  strings.ToUpper(record[5]),
		}
		cities = append(cities, city)
	}


	// Now 'cities' contains the data in a structured format
	// fmt.Println(cities)

	var n int  // distributors count YELLANDU-TELANGANA-INDIA

	fmt.Println("the program will assign permissions to distributor1 randomly on city basis in INDIA and distributor1 is head.")
	fmt.Println("so, all sub distributors will recieve permissions from him randomly")
    fmt.Print("Enter Number of Distributors: ")


    fmt.Scanln(&n)


	var distributors []Distributor

	myset := make(map[int]bool)

	var head Distributor

	for i:=0;i<20; {
		ind := rand.Intn(len(cities))

		if ind == 0 {
			continue;
		}

		var area Area

		if _,ok := myset[ind]; !ok {
		
			area.city = cities[ind].CityName
			area.state = cities[ind].ProvinceName
			area.country = cities[ind].CountryName
			myset[ind] = true
			i++

			flag := rand.Intn(2)

			if flag == 0 {
				head.include = append(head.include,area)
			}else{
				head.exclude = append(head.exclude,area)
			}
		}
	}

	distributors = append(distributors,head)

	for i:=1;i<n;i++{
		var distributor Distributor
		set1 := make(map[int]bool)
		set2 := make(map[int]bool)
		j := 0
		
		length := len(distributors[i-1].include) + len(distributors[i-1].exclude)

		j = rand.Intn(length)
		// YAMUNANAGAR-HARYANA-INDIA
		for k:=0;k<j; {

			flag := rand.Intn(2)
	
			var area Area

			// YELLAPUR-KARNATAKA-INDIA
			
	
			if flag == 0 {
				randomInt := rand.Intn(len(distributors[i-1].include))
				if _,ok := set1[randomInt]; !ok {
					area.city = distributors[i-1].include[randomInt].city
					area.state = distributors[i-1].include[randomInt].state
					area.country = distributors[i-1].include[randomInt].country
					distributor.include = append(distributor.include,area)
					set1[randomInt] = true
					k++
				}
			}else{
				randomInt := rand.Intn(len(distributors[i-1].exclude))
				if _,ok := set2[randomInt]; !ok {
					area.city = distributors[i-1].exclude[randomInt].city
					area.state = distributors[i-1].exclude[randomInt].state
					area.country = distributors[i-1].exclude[randomInt].country
					distributor.exclude = append(distributor.exclude,area)
					set2[randomInt] = true
					k++
				}
			}
			
		}
	
		distributors = append(distributors,distributor)
	}

	for i,d := range distributors {
		fmt.Println("included for distributor",i+1)
		fmt.Println(d.include)
		fmt.Println("excluded for distributor",i+1)
		fmt.Println(d.exclude)
		fmt.Println(" ")
	}

	var q int

	fmt.Print("Enter Number of queries: ")
	fmt.Scanln(&q)


	fmt.Println("For each query the program will say whether given query is true or not ")
	fmt.Println("please make sure given distributor falls in distributors range ")

	// WOKHA NAGALAND INDIA
	// distributor1-exclude-YADGIRI-KARNATAKA-INDIA

	for i:=0;i<q;i++ {
		fmt.Print("Enter Query in (distributor1-exclude-bangalore-karnataka-india) format: ")
		var query string
		
		reader := bufio.NewReader(os.Stdin)

		query,_ = reader.ReadString('\n')

		// fmt.Print(query)

		slicedList := strings.Split(query, "-")

		var distributorInd int

		dist := slicedList[0]
		
		lastChar := dist[len(dist)-1]

		distributorInd, err := strconv.Atoi(string(lastChar))

		if err != nil {
			fmt.Println("Error converting last character to integer to get the distributor id:", err)
			continue
		}

		if distributorInd > n {
			fmt.Println("given distributor id is greater than current distributors count")
			continue
		}

		if strings.ToUpper(slicedList[1]) == "EXCLUDE" {
			flag := isExcluded(distributorInd-1,distributors,slicedList[2],slicedList[3],slicedList[4])
			if  flag == 0 {
				fmt.Println("NO, This city is Included for this distributor")
			}else if flag == 1{
				fmt.Println("YES, This city is Excluded for this distributor")
			}else{
				fmt.Println("This city is not under this distributor region !")
			}
		}else{
			flag := isIncluded(distributorInd-1,distributors,slicedList[2],slicedList[3],slicedList[4])
			if  flag == 0 {
				fmt.Println("NO, This city is Excluded for this distributor")
			}else if flag == 1{
				fmt.Println("YES, This city is Included for this distributor")
			}else{
				fmt.Println("This city is not under this distributor region !")
			}
		}
	}

}

func isExcluded(ind int,distributors []Distributor,city,state,country string) int {
	for _,distributor := range distributors[ind].include {
		if strings.EqualFold(distributor.city,city) {
			return 0
		}
	}

	flag := false

	for _,distributor := range distributors[ind].exclude {
		if strings.EqualFold(distributor.city,city) {
			flag = true
		}
	}

	if flag == true {
		return 1
	}

	return 2
}

// distributor4-include-YELLANDU-TELANGANA-INDIA
func isIncluded(ind int,distributors []Distributor,city,state,country string) int {
	for _,distributor := range distributors[ind].exclude {
		if strings.EqualFold(distributor.city,city) && strings.EqualFold(distributor.state,state) {
			return 0
		}
	}

	flag := false
	for _,distributor := range distributors[ind].include {
		if strings.EqualFold(distributor.city,city) && strings.EqualFold(distributor.state,state) {
			flag = true
		}
	}

	if flag == true {
		return 1
	}

	return 2
}