package main

import (
	"bufio"
	"challenge.com/domain"
	"challenge.com/usecase"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	var mx sync.RWMutex
	domain.CountryMap = make(map[string][][]string)
	domain.DistributorMap = make(map[string]domain.Distributor)
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.File
	file, err := os.Open("cities.csv")

	// Checks for the error
	if err != nil {
		fmt.Println("Error while reading the file", err)
	}

	// Closes the file
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal("Error while closing the file", err)
		}
	}()

	// The csv.NewReader() function is called in
	// which the object os.File passed as its parameter
	// and this creates a new csv.Reader that reads
	// from the file
	reader := csv.NewReader(file)

	// ReadAll reads all the records from the CSV file
	// and Returns them as slice of slices of string
	// and an error if any
	records, err := reader.ReadAll()

	// Checks for the error
	if err != nil {
		fmt.Println("Error reading records")
	}

	// Loop to iterate through
	// and print each of the string slice
	for _, eachRecord := range records {
		domain.CountryMap[strings.ToLower(eachRecord[len(eachRecord)-1])] = append(domain.CountryMap[strings.ToLower(eachRecord[len(eachRecord)-1])], eachRecord)
	}

	flag := true
	for flag {
		var choice int
		fmt.Println("The options are :")
		fmt.Println("1.Create distributor and assign permissions")
		fmt.Println("2.Create Subdistributor and assign permissions")
		fmt.Println("3.Display Distributor details")
		fmt.Println("4.Delete Distributor")
		fmt.Println("5.Check Distributor Permissions")
		fmt.Println("6.Exit")
		_, err = fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error in getting choice")
		}
		switch choice {
		case 1:
			var includePermissions, excludePermissions []string
			var name string
			fmt.Println("enter name for distributor")
			_, err = fmt.Scanln(&name)
			if err != nil {
				fmt.Println("Error in getting name")
			}
			mx.RLock()
			_, ok := domain.DistributorMap[name]
			mx.RUnlock()
			if ok {
				fmt.Println("Distributor with this name already created")
			} else {
				includePermissions, excludePermissions = usecase.GetDistributorInput()
				err = usecase.CreateDistributor(name, includePermissions, excludePermissions, &mx)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		case 2:
			var includePermissions, excludePermissions []string
			var name string
			fmt.Println("enter name for sub distributor")
			_, err = fmt.Scanln(&name)
			if err != nil {
				fmt.Println("Error in getting sub distributor name")
			}
			var parentName string
			fmt.Println("enter name for Parent distributor")
			_, err = fmt.Scanln(&parentName)
			if err != nil {
				fmt.Println("Error in getting parent distributor name")
			}
			mx.RLock()
			_, pok := domain.DistributorMap[parentName]
			mx.RUnlock()
			if !pok {
				fmt.Println("ParentDistributor with this name does not exist")
			}
			mx.RLock()
			_, ok := domain.DistributorMap[name]
			mx.RUnlock()
			if ok {
				fmt.Println("Distributor with this name already created")
			}
			if pok && !ok {
				includePermissions, excludePermissions = usecase.GetDistributorInput()
				err = usecase.CreateSubDistributor(name, parentName, includePermissions, excludePermissions, &mx)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		case 3:
			fmt.Println("Give the distributor name")
			var name string
			_, err = fmt.Scanln(&name)
			if err != nil {
				fmt.Println("Error in getting choice")
			}
			usecase.GetDistributorDetails(name, &mx)
		case 4:
			fmt.Println("Give the distributor name")
			var name string
			_, err = fmt.Scanln(&name)
			if err != nil {
				fmt.Println("Error in getting choice")
			}
			usecase.DeleteDistributorDetails(name, &mx)
		case 5:
			fmt.Println("Give the distributor name")
			var name string
			_, err = fmt.Scanln(&name)
			if err != nil {
				fmt.Println("Error in getting name")
			}
			fmt.Println("Give area of format 'CHENNAI-TAMIL NADU-INDIA'  - is mandatory")
			fmt.Println("Enter Area:")
			var area string
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			err := scanner.Err()
			if err != nil {
				fmt.Println("Error in getting area")
			} else {
				area = scanner.Text()
			}
			permission, include, exclude, err := usecase.CheckUserPermissions(area, name, &mx)
			if err == nil {
				if permission {
					fmt.Println("Area is permitted for distributor ", name)
				} else {
					fmt.Println("Area is not permitted for distributor ", name)
				}
				if len(include) > 0 {
					fmt.Println("Included areas---", include)
				}
				if len(exclude) > 0 {
					fmt.Println("Excluded areas---", exclude)
				}
			}
			fmt.Println()
			fmt.Println()
		case 6:
			flag = false
		}
	}
}
