package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type CityData struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}

type Distributer struct {
	Code       string
	Name       string
	Desciption string
	City       string
	Includes   []string
	Excludes   []string
}

var CityMasterData map[string]CityData
var DistributerData map[string]Distributer
var StateMasterMap map[string][]string
var CountryMasterMap map[string][]string

func main() {
	fmt.Println("Welcome to Real Image Challenge 2016")
	fmt.Println("Seeding master data - started & processing...")
	InitData()
	fmt.Println("Seeding master data - completed...")
	fmt.Println("MasterData Len: ", len(CityMasterData))
SystemOption:
	fmt.Println("----------------------------------------------------")
	fmt.Println("System Option \n 1.Add Distibuter \n 2.Add Permission \n 3.Permission Test \n 4.Exit ")
	fmt.Println("----------------------------------------------------")
	arg := 0
	fmt.Scan(&arg)
	switch arg {
	case 1:
		fmt.Println("Enter Name: ")
		name := ""
		fmt.Scan(&name)
		fmt.Println("Enter City: ")
		city := ""
		fmt.Scan(&city)
		data, ok := DistributerData[strings.ToLower(name)]
		if ok {
			fmt.Println("Distributer already exit as ", data.Name, " in ", data.City, " city.")
		} else {
			DistributerData[strings.ToLower(name)] = Distributer{
				Name: name,
				Code: strings.ToLower(name),
				City: city,
			}
		}
		goto SystemOption
	case 2:
		fmt.Println("Enter Distributer Name: ")
		name := ""
		fmt.Scan(&name)
		data, ok := DistributerData[strings.ToLower(name)]
		if ok {
		pmnOpt:
			fmt.Println("Add Permissions: (1).Include rights,(2)Exclude rights, (3)Main Menu.")
			permissions := 0
			fmt.Scan(&permissions)
			switch permissions {
			case 1:
				permCity := ""
				fmt.Scan(&permCity)
				data.Includes = append(data.Includes, strings.ToUpper(strings.ReplaceAll(permCity, " ", "")))
				goto pmnOpt
			case 2:
				permCity := ""
				fmt.Scan(&permCity)
				data.Excludes = append(data.Excludes, strings.ToUpper(strings.ReplaceAll(permCity, " ", "")))
				goto pmnOpt
			case 3:
				goto SystemOption
			default:
				fmt.Println("enter valid option")
				goto pmnOpt
			}
		} else {
			fmt.Println("Distributer not found.")
		}
		goto SystemOption

	case 3:

		fmt.Println("Enter Distributer Name: ")
		name := ""
		fmt.Scan(&name)
		data, ok := DistributerData[strings.ToLower(name)]
		if ok {
			permCity := ""
			fmt.Println("Enter City text: ")
			fmt.Scan(&permCity)
			permCity = strings.ToUpper(strings.ReplaceAll(permCity, " ", ""))
			IsCity := false
			IsState := false
			IsCountry := false
			
			State := ""
			Country := ""
			splitArr := strings.Split(permCity, "-")
			switch len(splitArr) {
			case 1:
				IsCountry = true
				Country = splitArr[0]
			case 2:
				IsCountry = true
				IsState = true

				State = splitArr[0]
				Country = splitArr[1]
			case 3:
				IsCountry = true
				IsState = true
				IsCity = true

				State = splitArr[1]
				Country = splitArr[2]
			}

			result := ""
			for _, d := range data.Includes {
				if IsCountry && Country == d {
					result = "Yes"
				}
				if IsState && (State+"-"+Country == d) {
					result = "Yes"
				}
				if IsCity && permCity == d {
					result = "Yes"
				}
			}
			for _, d := range data.Excludes {
				if result == "Yes" {
					if IsCountry && Country == d {
						result = "No"
					}
					if IsState && (State+"-"+Country == d) {
						result = "No"
					}
					if IsCity && permCity == d {
						result = "No"
					}
				} else {
					if IsCountry && Country == d {
						result = "No"
					}
					if IsState && (State+"-"+Country == d) {
						result = "No"
					}
					if IsCity  && permCity == d {
						result = "No"
					}
				}
			}
			fmt.Println("Result: ", result)
		} else {
			fmt.Println("Distributer not found.")
		}
		goto SystemOption
	case 4:
		os.Exit(0)
	default:
		fmt.Println("enter valid option")
		goto SystemOption
	}
}

func InitData() {
	if CityMasterData == nil {
		CityMasterData = map[string]CityData{}
	}
	StateMasterMap = map[string][]string{}
	CountryMasterMap = map[string][]string{}

	DistributerData = map[string]Distributer{}
	DistributerData[strings.ToLower("Distributer1")] = Distributer{
		Name:     "Distributer1",
		City:     "Chennai",
		Includes: []string{"INDIA", "UNITEDSTATES"},
		Excludes: []string{"KARNATAKA-INDIA", "CHENNAI-TAMILNADU-INDIA"},
	}

	file, fErr := os.Open("cities.csv")
	defer file.Close()
	if fErr != nil {
		fmt.Println("File not found...", fErr.Error())
	}

	csvReader := csv.NewReader(file)
	data, derr := csvReader.ReadAll()

	if derr != nil {
		fmt.Println("data reading error...", derr.Error())
	}

	for i := range data {
		// fmt.Println(data[i])
		if i > 0 {
			indata := data[i]
			key := strings.ToLower(indata[0] + indata[1] + indata[2])
			CityMasterData[key] = CityData{
				CityCode:     indata[0],
				ProvinceCode: indata[1],
				CountryCode:  indata[2],
				CityName:     indata[3],
				ProvinceName: indata[4],
				CountryName:  indata[5],
			}
			CityName := strings.ToUpper(strings.ReplaceAll(indata[3], " ", ""))
			ProvinceName := strings.ToUpper(strings.ReplaceAll(indata[4], " ", ""))
			CountryName := strings.ToUpper(strings.ReplaceAll(indata[5], " ", ""))
			if cd, ok := CountryMasterMap[CountryName]; ok {
				idx := findIndex(cd, ProvinceName)
				if idx < 0 {
					cd = append(cd, ProvinceName)
				}
				CountryMasterMap[CountryName] = cd
			} else {
				CountryMasterMap[CountryName] = []string{ProvinceName}
			}

			if cd, ok := StateMasterMap[ProvinceName]; ok {
				idx := findIndex(cd, CityName)
				if idx < 0 {
					cd = append(cd, CityName)
				}
				StateMasterMap[ProvinceName] = cd
			} else {
				StateMasterMap[ProvinceName] = []string{CityName}
			}
		}
	}
}

func findIndex(slice []string, value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}
