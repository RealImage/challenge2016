package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/atyagi9006/challenge2016/csvreader"
	"github.com/atyagi9006/challenge2016/distributer"
	"github.com/atyagi9006/challenge2016/models"
	"github.com/atyagi9006/challenge2016/utilites"
)

func main() {
	csvFileName := "cities.csv"
	distributerMap := make(models.DistributerMap)
	countryStateMap := make(models.CountryMap)

	err := csvreader.MakeDataStore(csvFileName, countryStateMap)
	if err != nil {
		log.Fatalf("error making Data Store: %v", err)
	}

	DynamicInput(countryStateMap, distributerMap)
	mapJSON, _ := json.Marshal(distributerMap)
	fmt.Println("OutPut : " + string(mapJSON))
}

func DynamicInput(countryMap models.CountryMap, distributerMap models.DistributerMap) {
	inputSlice := make([]models.InputModel, 3)
	//need to implement ho to take input from console
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("press 0 to exit")
		fmt.Printf("enter distrbuter name :\t")
		scanner.Scan()
		distributerInput := utilites.UpperCaseNoSpace(scanner.Text())

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
					_, err := distributer.AddDistributer(input, countryMap, distributerMap)
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
					_, err := distributer.AddDistributer(input, countryMap, distributerMap)
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

	sliceJSON, _ := json.Marshal(inputSlice)
	fmt.Println("INPUT : " + string(sliceJSON))
}
