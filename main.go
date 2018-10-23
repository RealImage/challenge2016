package main

import (
	
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"github.com/atyagi9006/challenge2016/models"
	"github.com/atyagi9006/challenge2016/csvreader"
	"github.com/atyagi9006/challenge2016/distributer"
	"github.com/atyagi9006/challenge2016/utilites"
)

func main() {
	csvFileName := "cities.csv"
	distributerMap := make(models.DistributerMap)
	countryStateMap := make(models.CountryMap)

	csvreader.MakeDataStore(csvFileName, countryStateMap)
	
	//StaticInput(countryStateMap, distributerMap)
	DynamicInput(countryStateMap, distributerMap)
	mapJSON, _ := json.Marshal(distributerMap)
	fmt.Println("OutPut : "+string(mapJSON))
}

func StaticInput(countryMap models.CountryMap, distributerMap models.DistributerMap) {
	input := models.InputModel{
		Name:       utilites.UpperCaseNoSpace("distributer"),
		Permission: "India",
		AuthType:   models.Include,
	}
	distributer.AddDistributer(input, countryMap, distributerMap)

	input1 := models.InputModel{
		Name:       utilites.UpperCaseNoSpace("distributer"),
		Permission: "Tamil Nadu-India",
		AuthType:   models.Exclude,
	}
	distributer.AddDistributer(input1, countryMap, distributerMap)

	input2 := models.InputModel{
		Name:       utilites.UpperCaseNoSpace("distributer1 < distributer"),
		Permission: "Keelakarai-Tamil Nadu-India",
		AuthType:   models.Exclude,
	}
	distributer.AddDistributer(input2, countryMap, distributerMap) 
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
		}else if distributerInput != "" {
			fmt.Printf("Permissions for : %s \n", distributerInput)
			for {
				fmt.Printf("INCLUDE:")
				scanner.Scan()
				csc := scanner.Text()
				if csc == "0" {
					break
				}else if csc != "" {
					input := models.InputModel{
						Name:       distributerInput,
						Permission: csc,
						AuthType:   models.Include,
					}
					distributer.AddDistributer(input, countryMap, distributerMap)
					inputSlice = append(inputSlice, input)
				}else {
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
				}else if csc != "" {
					input := models.InputModel{
						Name:       distributerInput,
						Permission: csc,
						AuthType:   models.Exclude,
					}
					distributer.AddDistributer(input, countryMap, distributerMap)
					inputSlice = append(inputSlice, input)
				}else {
					fmt.Println("enter a valid include permission")
				}
				
				fmt.Printf("\n")
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
	
		}else{
			fmt.Println("enter a valid distributer")
		}
		}
		

	sliceJSON, _ := json.Marshal(inputSlice)
	fmt.Println("INPUT : "+string(sliceJSON))
}

