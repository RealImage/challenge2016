package main

import (
	"bufio"
	"example/controller"
	"example/helpers"
	"example/models"

	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvFileName := "cities.csv"
	distributerMap := make(models.DistributerMap)
	countryStateCityMap := make(models.CountryMap)

	err := helpers.CSVDataFetch(csvFileName, countryStateCityMap)
	if err != nil {
		log.Fatalf("error making Data Store: %v", err)
	}

	ConvertCSVDatatoMap(countryStateCityMap, distributerMap)

	var distributers []string
	for k := range distributerMap {
		distributers = append(distributers, k)
	}

	for {
		fmt.Print("Select distrbuter name from :\t", distributers)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		distributerName := UpperCaseNoSpace(scanner.Text())
		if distributerName == "0" {
			break
		}
		fmt.Printf("Enter country or state or city name")
		scanner.Scan()
		reasonName := scanner.Text()
		result := strings.Split(reasonName, "-")
		fmt.Println("", result)
		ss := distributerMap[distributerName]
		_, ok := ss.IncludeMap[result[len(result)-1]]
		if ok {
			for i := 0; i < len(strings.Split(reasonName, "-")); i++ {
				subset := generateSubset(reasonName, i)
				if _, ok := ss.ExcludeMap[subset]; ok {
					fmt.Printf("%v do not have permission to make distribution in %v .\n", distributerName, result)
					return
				} else {
					continue
				}
			}
			fmt.Printf("%v  have permission to make distribution in %v .\n", distributerName, result)
		} else {
			fmt.Printf("%v do not have permission to make distribution in %v .\n", distributerName, result)
		}
	}
}

func ConvertCSVDatatoMap(countryMap models.CountryMap, distributerMap models.DistributerMap) {
	inputSlice := []models.InputModel{}
	//need to implement ho to take input from console
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("press 0 to exit")
		fmt.Printf("enter distrbuter name :\t")
		scanner.Scan()
		distributerInput := UpperCaseNoSpace(scanner.Text())

		if distributerInput == "0" {
			break
		} else if distributerInput != "" {
			fmt.Printf("Permissions for : %s \n", distributerInput)
			for {
				fmt.Printf("INCLUDE:")
				scanner.Scan()
				csc := scanner.Text()
				if csc == "0" {
					break
				} else if csc != "" {
					input := models.InputModel{
						Name:       distributerInput,
						Permission: csc,
						AuthType:   models.Include,
					}
					_, err := controller.AddDistributer(input, countryMap, distributerMap)
					if err != nil {
						fmt.Printf("Error : %v \n", err)
					}
					inputSlice = append(inputSlice, input)
				} else {
					fmt.Println("enter a valid include permission")
				}

				fmt.Printf("\n")
			}
			for {
				fmt.Printf("ExCLUDE:")
				scanner.Scan()
				csc := scanner.Text()

				if csc == "0" {
					break
				} else if csc != "" {
					input := models.InputModel{
						Name:       distributerInput,
						Permission: csc,
						AuthType:   models.Exclude,
					}
					_, err := controller.AddDistributer(input, countryMap, distributerMap)
					if err != nil {
						fmt.Printf("Error : %v \n", err)
					}
					inputSlice = append(inputSlice, input)
				} else {
					fmt.Println("enter a valid include permission")
				}

				fmt.Printf("\n")
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}

		} else {
			fmt.Println("enter a valid distributer")
		}
	}
}

func UpperCaseNoSpace(input string) string {
	return strings.ToUpper(strings.Replace(input, " ", "", -1))
}

func generateSubset(input string, startIndex int) string {
	parts := strings.Split(input, "-")

	// Check if the startIndex is within bounds
	if startIndex < 0 || startIndex >= len(parts) {
		return "Invalid startIndex"
	}

	// Create a subset starting from the specified index
	subset := strings.Join(parts[startIndex:], "-")
	return subset
}
